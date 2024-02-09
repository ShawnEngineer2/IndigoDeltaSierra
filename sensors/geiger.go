package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func GeigerInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Geiger Counter sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Geiger.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Geiger.GeigerCounterState = 1
		(*qubzMatrix)[i].Geiger.GeigerReading = 10.649387
	}

}
