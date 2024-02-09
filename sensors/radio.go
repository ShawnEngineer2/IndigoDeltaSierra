package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func RadioInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Radio sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Radio.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Radio.GPSHandshakeStable = 1
		(*qubzMatrix)[i].Radio.GPSSignalStrength = -101
		(*qubzMatrix)[i].Radio.LocalWiFiHandshake = 1
		(*qubzMatrix)[i].Radio.LocalWiFiSignalStrength = -17.29
		(*qubzMatrix)[i].Radio.RadioFirmwareState = 1
		(*qubzMatrix)[i].Radio.RadioFirmwareVersion = "Motorola 6480 - V:17.2.5"
		(*qubzMatrix)[i].Radio.RadioPowerState = 1
		(*qubzMatrix)[i].Radio.RadioSensorState = 1
		(*qubzMatrix)[i].Radio.RadioSignalStrength = -85
		(*qubzMatrix)[i].Radio.WiFiDownlinkState = 1
		(*qubzMatrix)[i].Radio.WiFiDownloadSpeed = 940
		(*qubzMatrix)[i].Radio.WiFiUplinkState = 1
		(*qubzMatrix)[i].Radio.WiFiUploadSpeed = 11
	}

}
