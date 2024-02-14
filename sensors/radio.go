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

func RadioGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.RadioEvent {

	//Create Event Instance
	eventInstance := datamodels.RadioEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_RADIO
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_RADIO
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.RadioReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Radio.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].GPSHandshakeStable = (*qubzMatrixCurrent)[matrixIndex].Radio.GPSHandshakeStable
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].GPSSignalStrength = (*qubzMatrixCurrent)[matrixIndex].Radio.GPSSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LocalWiFiHandshake = (*qubzMatrixCurrent)[matrixIndex].Radio.LocalWiFiHandshake
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LocalWiFiSignalStrength = (*qubzMatrixCurrent)[matrixIndex].Radio.LocalWiFiSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RadioFirmwareState = (*qubzMatrixCurrent)[matrixIndex].Radio.RadioFirmwareState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RadioFirmwareVersion = (*qubzMatrixCurrent)[matrixIndex].Radio.RadioFirmwareVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RadioPowerState = (*qubzMatrixCurrent)[matrixIndex].Radio.RadioPowerState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RadioSensorState = (*qubzMatrixCurrent)[matrixIndex].Radio.RadioSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RadioSignalStrength = (*qubzMatrixCurrent)[matrixIndex].Radio.RadioSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].WiFiDownlinkState = (*qubzMatrixCurrent)[matrixIndex].Radio.WiFiDownlinkState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].WiFiDownloadSpeed = (*qubzMatrixCurrent)[matrixIndex].Radio.WiFiDownloadSpeed
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].WiFiUplinkState = (*qubzMatrixCurrent)[matrixIndex].Radio.WiFiUplinkState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].WiFiUploadSpeed = (*qubzMatrixCurrent)[matrixIndex].Radio.WiFiUploadSpeed

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Radio.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].GPSHandshakeStable = (*qubzMatrixPrevious)[matrixIndex].Radio.GPSHandshakeStable
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].GPSSignalStrength = (*qubzMatrixPrevious)[matrixIndex].Radio.GPSSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LocalWiFiHandshake = (*qubzMatrixPrevious)[matrixIndex].Radio.LocalWiFiHandshake
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LocalWiFiSignalStrength = (*qubzMatrixPrevious)[matrixIndex].Radio.LocalWiFiSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RadioFirmwareState = (*qubzMatrixPrevious)[matrixIndex].Radio.RadioFirmwareState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RadioFirmwareVersion = (*qubzMatrixPrevious)[matrixIndex].Radio.RadioFirmwareVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RadioPowerState = (*qubzMatrixPrevious)[matrixIndex].Radio.RadioPowerState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RadioSensorState = (*qubzMatrixPrevious)[matrixIndex].Radio.RadioSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RadioSignalStrength = (*qubzMatrixPrevious)[matrixIndex].Radio.RadioSignalStrength
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].WiFiDownlinkState = (*qubzMatrixPrevious)[matrixIndex].Radio.WiFiDownlinkState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].WiFiDownloadSpeed = (*qubzMatrixPrevious)[matrixIndex].Radio.WiFiDownloadSpeed
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].WiFiUplinkState = (*qubzMatrixPrevious)[matrixIndex].Radio.WiFiUplinkState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].WiFiUploadSpeed = (*qubzMatrixPrevious)[matrixIndex].Radio.WiFiUploadSpeed

	//Return the completed event
	return eventInstance

}
