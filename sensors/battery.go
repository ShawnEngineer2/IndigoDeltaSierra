package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func BatteryInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Battery sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Battery.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Battery.BatteryChargePct = 100
		(*qubzMatrix)[i].Battery.BatterySensorState = 1
		(*qubzMatrix)[i].Battery.BatteryTemperature = 35
		(*qubzMatrix)[i].Battery.CapacitanceGelActualVolume = 136
		(*qubzMatrix)[i].Battery.CapacitanceGelExpectedVolume = 136
		(*qubzMatrix)[i].Battery.DrainRate = 450
		(*qubzMatrix)[i].Battery.MaxAmpHours = 4780
		(*qubzMatrix)[i].Battery.RemainingAmpHours = 4290

	}

}
