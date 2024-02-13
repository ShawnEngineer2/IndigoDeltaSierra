package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func RadioInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Radio sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Radio.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Radio.GPSHandshakeStable = 1
	(*qubzMatrix)[matrixIndex].Radio.GPSSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_GPS_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.LocalWiFiHandshake = 1
	(*qubzMatrix)[matrixIndex].Radio.LocalWiFiSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_LOCAL_WIFI_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.RadioFirmwareState = 1
	(*qubzMatrix)[matrixIndex].Radio.RadioFirmwareVersion = "Motorola 6480 - V:17.2.5"
	(*qubzMatrix)[matrixIndex].Radio.RadioPowerState = 1
	(*qubzMatrix)[matrixIndex].Radio.RadioSensorState = 1
	(*qubzMatrix)[matrixIndex].Radio.RadioSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_RADIO_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.WiFiDownlinkState = 1
	(*qubzMatrix)[matrixIndex].Radio.WiFiDownloadSpeed = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_DOWNLOAD_SPEED, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.WiFiUplinkState = 1
	(*qubzMatrix)[matrixIndex].Radio.WiFiUploadSpeed = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_UPLOAD_SPEED, qubzStateDS, consoleLogger, fileLogger)

}
