package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func LockInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Lock sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Lock.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Lock.LockEventTime = "02/08/2024 21:37:10.093"
	(*qubzMatrix)[matrixIndex].Lock.LockEventMethod = 3
	(*qubzMatrix)[matrixIndex].Lock.LockEventType = 1
	(*qubzMatrix)[matrixIndex].Lock.LockState = 1

}
