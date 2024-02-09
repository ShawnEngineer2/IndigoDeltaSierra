package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func GPSInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the GPS sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].GPS.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].GPS.GPSSensorState = 1
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Latitude.Hours = 32
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Latitude.Minutes = 47
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Latitude.Seconds = 8
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Longitude.Hours = -116
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Longitude.Minutes = 2
		(*qubzMatrix)[i].GPS.ActualGPSPosition.Longitude.Seconds = 58

	}

}
