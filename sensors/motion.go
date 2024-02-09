package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func MotionInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Motion sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Motion.MotionSensorState = 1
		(*qubzMatrix)[i].Motion.NumberOfContacts = 0
		(*qubzMatrix)[i].Motion.AverageVelocity = 0

	}

}
