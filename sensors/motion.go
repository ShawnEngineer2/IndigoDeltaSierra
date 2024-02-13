package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func MotionInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Motion sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Motion.MotionSensorState = 1
	(*qubzMatrix)[matrixIndex].Motion.NumberOfContacts = 0
	(*qubzMatrix)[matrixIndex].Motion.AverageVelocity = 0

}