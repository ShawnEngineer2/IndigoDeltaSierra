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

func createSensorState(qubzMatrix datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, fileLogger *slog.Logger) (datamodels.QubzState, error) {

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
			sensorState.SensorDataPoints[i].DataPointValue = generateRangedSensorValue(x.NominalMin, x.NominalMax, x.NumberScale, fileLogger)

		case appconstants.DATA_POINT_TYPE_BOOLEAN:
			sensorState.SensorDataPoints[i].DataPointValue = x.NominalMin

		default:
			sensorState.SensorDataPoints[i].DataPointValue = 0.00
		}

	}

	return sensorState, nil
}

func generateRangedSensorValue(minValue float64, maxValue float64, numScale int, fileLogger *slog.Logger) float64 {

	//This routine generates a random value with the range specified. It then ensures the returned
	//values are within the requested range as a safety precaution
	newValue := randomgen.RandomIFloat(minValue, maxValue, numScale)

	if newValue < minValue {
		newValue = minValue
		customlog.InfoFile(fileLogger, "Trim Requested Range up to requested Min Value")

	} else if newValue > maxValue {
		newValue = maxValue
		customlog.InfoFile(fileLogger, "Trim Requested Range up to requested Max Value")

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

	//If this Qubz Unit has an exception with an Interval Boundary, adjust and advance the boundary before setting sensor data point values
	if sensorState.ExceptionType != appconstants.SENSOR_EXCEPTION_TYPE_NONE && sensorState.ExceptionIntervalBoundary > 0 {
		//Adjust and advance the interval counter as needed
		if sensorState.CurrentExceptionInterval == sensorState.ExceptionIntervalBoundary {
			//Restart the intermittency interval count
			sensorState.CurrentExceptionInterval = 1
		} else {
			//Advance the intermittency interval counter
			sensorState.CurrentExceptionInterval++
		}
	}

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
			sensorState.SensorDataPoints[i].DataPointValue = handleSensorNominalState(x, fileLogger)

		} else {
			//Handle intermittency (error value generated intermittently)
			var newValue float64

			if sensorState.ExceptionIntervalBoundary > 0 {
				//This is an intermittent exception - determine if generate an Exception value or a Nominal value
				if sensorState.CurrentExceptionInterval == sensorState.ExceptionIntervalBoundary {
					//Generate an exception value
					newValue = handleSensorExceptionState(x, qubzException, fileLogger)
				} else {
					//Generate a nominal value
					newValue = handleSensorNominalState(x, fileLogger)
				}

			} else {
				newValue = handleSensorExceptionState(x, qubzException, fileLogger)
			}

			//Acquire the correct value for this sensor data point
			sensorState.SensorDataPoints[i].DataPointValue = newValue

		}

		//fmt.Println(sensorState.SensorDataPoints[i].DataPointValue)

	}

	return sensorState, nil
}

func handleSensorNominalState(sensorRange datamodels.SensorRange, fileLogger *slog.Logger) float64 {

	//Generate and return a value in the nominal range
	var newValue float64

	switch sensorRange.DataPointTypeId {
	case appconstants.DATA_POINT_TYPE_RANGED:
		//customlog.CalloutConsole(consoleLogger, "In Range Value Block")
		newValue = generateRangedSensorValue(sensorRange.NominalMin, sensorRange.NominalMax, sensorRange.NumberScale, fileLogger)

	case appconstants.DATA_POINT_TYPE_BOOLEAN:
		newValue = sensorRange.NominalMin

	}

	return newValue
}

func handleSensorExceptionState(sensorRange datamodels.SensorRange, exceptionDef datamodels.QubzException, fileLogger *slog.Logger) float64 {

	//Walk through the sensor data points in the Exception definition and if one of them matches the passed sensor data point id
	//then set a value based on Exception Type

	var newValue float64

	for _, x := range exceptionDef.AffectedSensors {

		if x.SensorDataPointId == sensorRange.SensorDataPointId {

			//Set an exception value based on the value mod type of the Sensor Exception
			switch x.ValueModType {
			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_FIXED:
				newValue = x.FixedModValue

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HARD_HIGH:
				newValue = sensorRange.ExceptionMax

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HARD_LOW:
				newValue = sensorRange.ExceptionMin

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_HIGH:
				//Determine the appropriate minimum value and then calculate
				var newMin float64

				if sensorRange.ExceptionMin > sensorRange.NominalMax {
					newMin = sensorRange.ExceptionMin
				} else {
					newMin = sensorRange.NominalMax
				}

				newValue = randomgen.RandomIFloat(newMin, sensorRange.ExceptionMax, sensorRange.NumberScale)

			case appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW:
				//Determine the appropriate Maximum value and then calculate
				var newMax float64

				if sensorRange.ExceptionMax < sensorRange.NominalMin {
					newMax = sensorRange.ExceptionMax
				} else {
					newMax = sensorRange.NominalMin
				}

				newValue = randomgen.RandomIFloat(sensorRange.ExceptionMin, newMax, sensorRange.NumberScale)

			}

			//Return the new values
			return newValue
		}
	}

	//If you get this far then there was no match on the exception data point - just calculate as a normal sensor reading
	newValue = handleSensorNominalState(sensorRange, fileLogger)

	return newValue
}
