package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func LockInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Lock sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Lock.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Lock.LockEventTime = "02/08/2024 21:37:10.093"
		(*qubzMatrix)[i].Lock.LockEventMethod = 3
		(*qubzMatrix)[i].Lock.LockEventType = 1
		(*qubzMatrix)[i].Lock.LockState = 1
	}

}
