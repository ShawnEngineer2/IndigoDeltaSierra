package simulator

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/randomgen"
	"log/slog"
)

//This package has routines to manage the creation and retrieval of new sensor states for qubz

func CreateSensorState(qubzMatrix datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, exceptionDS *[]datamodels.QubzException, consoleLogger *slog.Logger, fileLogger *slog.Logger) (datamodels.QubzState, error) {

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

	//Generate nominal values for sensors that are either ranged or boolean. Don't worry about Info or Calculated as
	//these are set in individual sensor update routines.

	for i, x := range *sensorRangeDS {

		//Fill in info
		sensorState.SensorDataPoints[i].DataPointTypeId = x.DataPointTypeId
		sensorState.SensorDataPoints[i].SensorDataPointId = x.SensorDataPointId

		//Calculate the value based on the type of data point
		switch x.DataPointTypeId {
		case appconstants.DATA_POINT_TYPE_RANGED:
			newValue, err := generateRangedSensorValue(x, exceptionDS, qubzMatrix.ExceptionAssignment, qubzMatrix.ExceptionType, qubzMatrix.CurrentExceptionInterval)

			if err != nil {
				return sensorState, err
			}

			sensorState.SensorDataPoints[i].DataPointValue = newValue

		case appconstants.DATA_POINT_TYPE_BOOLEAN:
			newValue, err := generateBooleanSensorValue(x, exceptionDS, qubzMatrix.ExceptionAssignment, qubzMatrix.ExceptionType, qubzMatrix.CurrentExceptionInterval)

			if err != nil {
				return sensorState, err
			}

			sensorState.SensorDataPoints[i].DataPointValue = newValue

		default:
			sensorState.SensorDataPoints[i].DataPointValue = 0.00
		}

	}

	return sensorState, nil
}

func generateRangedSensorValue(sensorInfo datamodels.SensorRange, exceptionDS *[]datamodels.QubzException, exceptionId int, exceptionType int, currentExceptionInterval int) (float64, error) {

	//This routine generates a new ranged value for the passed sensor, accomodating both nominal and exception states
	//as determined by the parameters passed into this routine
	var newValue float64 = 0
	var err error = nil

	if exceptionType == appconstants.SENSOR_EXCEPTION_TYPE_NONE {
		newValue = randomgen.RandomIFloat(sensorInfo.NominalMin, sensorInfo.NominalMax, sensorInfo.NumberScale)

	} else {
		newValue, err = generateSensorExceptionValue(sensorInfo, exceptionDS, exceptionId, exceptionType, currentExceptionInterval)

		if err != nil {
			return 0, err
		}
	}

	return newValue, nil
}

func generateBooleanSensorValue(sensorInfo datamodels.SensorRange, exceptionDS *[]datamodels.QubzException, exceptionId int, exceptionType int, currentExceptionInterval int) (float64, error) {

	//This routine generates a new boolean value for the passed sensor, accomodating both nominal and exception states
	//as determined by the parameters passed into this routine
	var newValue float64 = 0
	var err error = nil

	if exceptionType == appconstants.SENSOR_EXCEPTION_TYPE_NONE {
		newValue = sensorInfo.NominalMin
	} else {
		newValue, err = generateSensorExceptionValue(sensorInfo, exceptionDS, exceptionId, exceptionType, currentExceptionInterval)

		if err != nil {
			return 0, err
		}
	}

	return newValue, nil
}

func generateSensorExceptionValue(sensorInfo datamodels.SensorRange, exceptionDS *[]datamodels.QubzException, exceptionId int, exceptionType int, currentExceptionInterval int) (float64, error) {

	//Generate a value for the passed exception based on the information provided.
	//First, get the definition for the exception
	return 0, nil
}
