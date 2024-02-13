package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func FireInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Fire sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Fire.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Fire.FireDetected = 0
	(*qubzMatrix)[matrixIndex].Fire.FireSensorState = 1
	(*qubzMatrix)[matrixIndex].Fire.FireSuppressionState = 0

}
