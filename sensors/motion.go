package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func MotionInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Motion sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Motion.MotionSensorState = 1
	(*qubzMatrix)[matrixIndex].Motion.NumberOfContacts = 0
	(*qubzMatrix)[matrixIndex].Motion.AverageVelocity = 0

}

func MotionSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Motion sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Motion.MotionSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_MOTIONSENSORSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Motion.NumberOfContacts = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_NUMBEROFCONTACTS, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Motion.AverageVelocity = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_AVERAGEVELOCITY, qubzStateDS, consoleLogger, fileLogger)

}
