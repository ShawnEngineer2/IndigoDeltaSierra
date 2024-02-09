package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func QubzSealInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Qubz Seal sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].QubzSeal.SealSensorState = 1
		(*qubzMatrix)[i].QubzSeal.SealState = 1
		(*qubzMatrix)[i].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

	}

}
