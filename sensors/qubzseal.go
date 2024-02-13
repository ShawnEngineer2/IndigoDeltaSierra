package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func QubzSealInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Qubz Seal sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].QubzSeal.SealSensorState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

}
