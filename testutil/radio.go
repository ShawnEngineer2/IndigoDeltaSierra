package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func RadioTestEvent(outputFilePath string) {
	//This routine is used to test the Radio JSON

	jsondata := datamodels.RadioEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.RadioReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].GPSHandshakeStable = 1
	jsondata.SensorData[0].GPSSignalStrength = 40
	jsondata.SensorData[0].LocalWiFiHandshake = 1
	jsondata.SensorData[0].LocalWiFiSignalStrength = 122
	jsondata.SensorData[0].RadioFirmwareState = 1
	jsondata.SensorData[0].RadioFirmwareVersion = "5XWG678R"
	jsondata.SensorData[0].RadioPowerState = 1
	jsondata.SensorData[0].RadioSensorState = 1
	jsondata.SensorData[0].RadioSignalStrength = 89
	jsondata.SensorData[0].WiFiDownlinkState = 1
	jsondata.SensorData[0].WiFiDownloadSpeed = 149
	jsondata.SensorData[0].WiFiUplinkState = 1
	jsondata.SensorData[0].WiFiUploadSpeed = 11

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].GPSHandshakeStable = 1
	jsondata.SensorData[1].GPSSignalStrength = 40
	jsondata.SensorData[1].LocalWiFiHandshake = 1
	jsondata.SensorData[1].LocalWiFiSignalStrength = 122
	jsondata.SensorData[1].RadioFirmwareState = 1
	jsondata.SensorData[1].RadioFirmwareVersion = "5XWG678R"
	jsondata.SensorData[1].RadioPowerState = 1
	jsondata.SensorData[1].RadioSensorState = 1
	jsondata.SensorData[1].RadioSignalStrength = 89
	jsondata.SensorData[1].WiFiDownlinkState = 1
	jsondata.SensorData[1].WiFiDownloadSpeed = 149
	jsondata.SensorData[1].WiFiUplinkState = 1
	jsondata.SensorData[1].WiFiUploadSpeed = 11

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
