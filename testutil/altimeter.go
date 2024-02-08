package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func AltimeterTestEvent(outputFilePath string) {
	//This routine is used to test the altimeter JSON

	jsondata := datamodels.AltimeterEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.SensorType.SensorTypeId = 1
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeDescription = "altimeter"

	jsondata.SensorData = make([]datamodels.AltimeterReading, 2)

	jsondata.SensorData[0].AltimeterState = 1
	jsondata.SensorData[0].Altitude = 2000.265
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].AltimeterState = 1
	jsondata.SensorData[1].Altitude = 2330.365
	jsondata.SensorData[1].EventState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
