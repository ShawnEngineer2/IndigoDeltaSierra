package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func TempBarometricInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Temperature and Barometric sensors

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].TempBarometrics.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].TempBarometrics.BarometricSensorState = 1
		(*qubzMatrix)[i].TempBarometrics.HumidityLevel = 62.89709
		(*qubzMatrix)[i].TempBarometrics.MoistureLevel = 2.00987
		(*qubzMatrix)[i].TempBarometrics.OverallInternalTemperature = 12
		(*qubzMatrix)[i].TempBarometrics.TemperatureSensorState = 1

	}

}
