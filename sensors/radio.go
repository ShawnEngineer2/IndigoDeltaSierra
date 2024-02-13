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

func RadioSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Radio sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Radio.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Radio.GPSHandshakeStable = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_GPS_HANDSHAKE_STABLE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.GPSSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_GPS_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.LocalWiFiHandshake = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_LOCAL_WIFI_HANDSHAKE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.LocalWiFiSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_LOCAL_WIFI_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.RadioFirmwareState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_RADIO_FIRMWARE_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.RadioFirmwareVersion = "Motorola 6480 - V:17.2.5"
	(*qubzMatrix)[matrixIndex].Radio.RadioPowerState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_RADIO_POWER_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.RadioSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_RADIO_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.RadioSignalStrength = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_RADIO_SIGNAL_STRENGTH, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.WiFiDownlinkState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_DOWNLINK_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.WiFiDownloadSpeed = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_DOWNLOAD_SPEED, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Radio.WiFiUplinkState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_UPLINK_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Radio.WiFiUploadSpeed = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_RADIO_WIFI_UPLOAD_SPEED, qubzStateDS, consoleLogger, fileLogger)

}
