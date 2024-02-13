package simulator

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/randomgen"
	"log/slog"
)

//This package has routines to manage the creation and retrieval of new sensor states for qubz

func CreateSensorState(qubzMatrix datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) (datamodels.QubzState, error) {

	//This routine creates a new sensor state and returns it By Value to the caller

	sensorState := datamodels.QubzState{}
	sensorState.SensorDataPoints = make([]datamodels.QubzSensorDataPoint, len(*sensorRangeDS))

	//Update header information for use by the Qubz state (probably got some unnecessarily redundant data here - come back and check in next version)
	sensorState.ExceptionAssignment = qubzMatrix.ExceptionAssignment
	sensorState.ExceptionIntermittent = qubzMatrix.ExceptionIntermittent
	sensorState.ExceptionInterval = qubzMatrix.ExceptionInterval
	sensorState.ExceptionSeverity = qubzMatrix.ExceptionSeverity
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
			sensorState.SensorDataPoints[i].DataPointValue = randomgen.RandomIFloat(x.NominalMin, x.NominalMax, x.NumberScale)

		case appconstants.DATA_POINT_TYPE_BOOLEAN:
			sensorState.SensorDataPoints[i].DataPointValue = x.NominalMin

		default:
			sensorState.SensorDataPoints[i].DataPointValue = 0.00
		}

	}

	return sensorState, nil
}
