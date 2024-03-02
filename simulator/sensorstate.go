package simulator

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"indigodeltasierra/randomgen"
	"log/slog"
)

//This package has routines to manage the creation and retrieval of new sensor states for qubz

func createSensorState(qubzMatrix datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) (datamodels.QubzState, error) {

	//This routine creates a new sensor state and returns it By Value to the caller

	sensorState := datamodels.QubzState{}
	sensorState.SensorDataPoints = make([]datamodels.QubzSensorDataPoint, len(*sensorRangeDS))

	//Update header information for use by the Qubz state (probably got some unnecessarily redundant data here - come back and check in next version)
	sensorState.ExceptionAssignment = qubzMatrix.ExceptionAssignment
	sensorState.ExceptionSeverity = qubzMatrix.ExceptionSeverity
	sensorState.ExceptionType = qubzMatrix.ExceptionType
	sensorState.ExceptionIntervalBoundary = qubzMatrix.ExceptionIntervalBoundary
	sensorState.CurrentExceptionInterval = qubzMatrix.CurrentExceptionInterval
	sensorState.QubzID = qubzMatrix.QubzID
	sensorState.QubzName = qubzMatrix.QubzName
	sensorState.RouteAssignment = qubzMatrix.RouteAssignment
	sensorState.ShipmentType = qubzMatrix.ShipmentType
	sensorState.TransportMode = qubzMatrix.TransportMode

	//Generate nominal values for sensors that are either ranged or boolean. Don't worry about Info or Calculated and
	//don't worry about exceptions. Those are set by other routines either in individual sensors or in other routines
	//in this package
	for i, x := range *sensorRangeDS {

		//Fill in info
		sensorState.SensorDataPoints[i].DataPointTypeId = x.DataPointTypeId
		sensorState.SensorDataPoints[i].SensorDataPointId = x.SensorDataPointId

		//Calculate the value based on the type of data point
		switch x.DataPointTypeId {
		case appconstants.DATA_POINT_TYPE_RANGED:
			sensorState.SensorDataPoints[i].DataPointValue = generateRangedSensorValue(x.NominalMin, x.NominalMax, x.NumberScale, consoleLogger, fileLogger)

		case appconstants.DATA_POINT_TYPE_BOOLEAN:
			sensorState.SensorDataPoints[i].DataPointValue = x.NominalMin

		default:
			sensorState.SensorDataPoints[i].DataPointValue = 0.00
		}

	}

	return sensorState, nil
}

func generateRangedSensorValue(minValue float64, maxValue float64, numScale int, consoleLogger *slog.Logger, fileLogger *slog.Logger) float64 {

	//This routine generates a random value with the range specified. It then ensures the returned
	//values are within the requested range as a safety precaution
	newValue := randomgen.RandomIFloat(minValue, maxValue, numScale)

	if newValue < minValue {
		newValue = minValue
		//customlog.InfoAllChannels(consoleLogger, fileLogger, "Trim Requested Range up to requested Min Value", false)

	} else if newValue > maxValue {
		newValue = maxValue
		//customlog.InfoAllChannels(consoleLogger, fileLogger, "Trim Requested Range up to requested Max Value", false)

	}

	//customlog.CalloutConsole(consoleLogger, fmt.Sprintf("New Range Value is %f", newValue))

	return newValue
}

func generateBooleanSensorValue(consoleLogger *slog.Logger, fileLogger *slog.Logger) float64 {

	//This routine generates a random boolean value. It then ensures the returned
	//values are within the requested range as a safety precaution
	newValue := randomgen.RandomBool()

	if newValue < 0 {
		newValue = 0
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Boolean Trimmed up to 0 (False)", true)

	} else if newValue > 1 {
		newValue = 1
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Boolean Trimmed down to 1 (True)", true)

	}

	return float64(newValue)
}

func updateSensorState(qubzMatrix datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, exceptionDS *[]datamodels.QubzException, consoleLogger *slog.Logger, fileLogger *slog.Logger) (datamodels.QubzState, error) {

	//This routine creates a new sensor state and returns it By Value to the caller. It walks through each sensor in the
	//passed Qubz unit and generates values based on the configuration of the Qubz unit

	//customlog.CalloutConsole(consoleLogger, "In Update Sensor Routine")

	sensorState := datamodels.QubzState{}
	sensorState.SensorDataPoints = make([]datamodels.QubzSensorDataPoint, len(*sensorRangeDS))

	//Update header information for use by the Qubz state (probably got some unnecessarily redundant data here - come back and check in next version)
	sensorState.ExceptionAssignment = qubzMatrix.ExceptionAssignment
	sensorState.ExceptionSeverity = qubzMatrix.ExceptionSeverity
	sensorState.ExceptionType = qubzMatrix.ExceptionType
	sensorState.ExceptionIntervalBoundary = qubzMatrix.ExceptionIntervalBoundary
	sensorState.CurrentExceptionInterval = qubzMatrix.CurrentExceptionInterval
	sensorState.QubzID = qubzMatrix.QubzID
	sensorState.QubzName = qubzMatrix.QubzName
	sensorState.RouteAssignment = qubzMatrix.RouteAssignment
	sensorState.ShipmentType = qubzMatrix.ShipmentType
	sensorState.TransportMode = qubzMatrix.TransportMode

	//Testing only
	//sensorState.ExceptionType = appconstants.SENSOR_EXCEPTION_TYPE_CONTINUOUS
	//sensorState.ExceptionAssignment = 3

	//If the Qubz unit has an exception state, get details for the excepton assigned to it
	var qubzException datamodels.QubzException
	var err error

	qubzException = datamodels.QubzException{}

	if sensorState.ExceptionType != appconstants.SENSOR_EXCEPTION_TYPE_NONE {
		qubzException, err = datautil.GetSingleException(exceptionDS, sensorState.ExceptionAssignment)

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Error Retrieving Exception for Qubz Unit %s : Exception Id %d : %s", sensorState.QubzName, sensorState.ExceptionAssignment, err.Error()))
			return sensorState, fmt.Errorf("qubz Unit %s not processed due to exception retrieval error", sensorState.QubzName)
		}
	}

	//fmt.Println(qubzException)

	//Generate new values for each sensor
	for i, x := range *sensorRangeDS {

		//Fill in info
		sensorState.SensorDataPoints[i].DataPointTypeId = x.DataPointTypeId
		sensorState.SensorDataPoints[i].SensorDataPointId = x.SensorDataPointId

		//If this Qubz unit has an exception state, send it to a different routine for handling - otherwise set new value here
		if sensorState.ExceptionType == appconstants.SENSOR_EXCEPTION_TYPE_NONE {

			//customlog.CalloutConsole(consoleLogger, "In Sensor Update Block")

			//Calculate the value based on the type of data point - note: No DEFAULT block on this one so that fixed values set in sensor init will persist from run to run
			switch x.DataPointTypeId {
			case appconstants.DATA_POINT_TYPE_RANGED:
				//customlog.CalloutConsole(consoleLogger, "In Range Value Block")
				sensorState.SensorDataPoints[i].DataPointValue = generateRangedSensorValue(x.NominalMin, x.NominalMax, x.NumberScale, consoleLogger, fileLogger)

			case appconstants.DATA_POINT_TYPE_BOOLEAN:
				sensorState.SensorDataPoints[i].DataPointValue = x.NominalMin

			}
		} else {
			//Acquire the correct value for this sensor data point
			newValue, newIntervalValue := handleSensorExceptionState(x, qubzException, sensorState.CurrentExceptionInterval, consoleLogger, fileLogger)

			//Assign the value to the data point
			sensorState.SensorDataPoints[i].DataPointValue = newValue
			sensorState.CurrentExceptionInterval = newIntervalValue

		}

		//fmt.Println(sensorState.SensorDataPoints[i].DataPointValue)

	}

	return sensorState, nil
}

func handleSensorExceptionState(sensorRange datamodels.SensorRange, exceptionDef datamodels.QubzException, currentExceptionInterval int, consoleLogger *slog.Logger, fileLogger *slog.Logger) (float64, int) {

	//Walk through the sensor data points in the Exception definition and if one of them matches the passed sensor data point id
	//then set a value based on Exception Type

	var newValue float64
	var newExceptionInterval int

	for _, x := range exceptionDef.AffectedSensors {

		if x.SensorDataPointId == sensorRange.SensorDataPointId {

			//Set an exception value based on the value mod type of the Sensor Exception
			switch x.ValueModType {
			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_FIXED:
				newValue = x.FixedModValue
				newExceptionInterval = currentExceptionInterval

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HARD_HIGH:
				newValue = sensorRange.ExceptionMax
				newExceptionInterval = currentExceptionInterval

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HARD_LOW:
				newValue = sensorRange.ExceptionMin
				newExceptionInterval = currentExceptionInterval

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HIGH:
				//Determine the appropriate minimum value and then calculate
				var newMin float64

				if sensorRange.ExceptionMin > sensorRange.NominalMax {
					newMin = sensorRange.ExceptionMin
				} else {
					newMin = sensorRange.NominalMax
				}

				newValue = randomgen.RandomIFloat(newMin, sensorRange.ExceptionMax, sensorRange.NumberScale)
				newExceptionInterval = currentExceptionInterval

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW:
				//Determine the appropriate Maximum value and then calculate
				var newMax float64

				if sensorRange.ExceptionMax < sensorRange.NominalMin {
					newMax = sensorRange.ExceptionMax
				} else {
					newMax = sensorRange.NominalMin
				}

				newValue = randomgen.RandomIFloat(sensorRange.ExceptionMin, newMax, sensorRange.NumberScale)
				newExceptionInterval = currentExceptionInterval

			}

			//Return the new values
			return newValue, newExceptionInterval
		}
	}

	//If you get this far then there was no match on the exception data point - just calculate as a normal sensor reading
	newExceptionInterval = currentExceptionInterval //Just set here since this isn't relevant

	switch sensorRange.DataPointTypeId {
	case appconstants.DATA_POINT_TYPE_RANGED:
		//customlog.CalloutConsole(consoleLogger, "In Range Value Block")
		newValue = generateRangedSensorValue(sensorRange.NominalMin, sensorRange.NominalMax, sensorRange.NumberScale, consoleLogger, fileLogger)

	case appconstants.DATA_POINT_TYPE_BOOLEAN:
		newValue = sensorRange.NominalMin

	}

	return newValue, newExceptionInterval
}
