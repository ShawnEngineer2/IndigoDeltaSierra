package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func FireInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Fire sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Fire.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Fire.FireDetected = 0
		(*qubzMatrix)[i].Fire.FireSensorState = 1
		(*qubzMatrix)[i].Fire.FireSuppressionState = 0

	}

}
