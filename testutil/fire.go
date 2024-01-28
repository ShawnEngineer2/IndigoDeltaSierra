package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func FireTestEvent(outputFilePath string) {
	//This routine is used to test the fire JSON

	jsondata := datamodels.FireEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.SensorType.SensorTypeId = 13
	jsondata.EventHeader.SensorType.SensorTypeDescription = "fire"

	jsondata.SensorData = make([]datamodels.FireReading, 2)

	jsondata.SensorData[0].FireDetected = 0
	jsondata.SensorData[0].FireSensorState = 1
	jsondata.SensorData[0].FireSuppressionState = 0
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].FireDetected = 0
	jsondata.SensorData[1].FireSensorState = 1
	jsondata.SensorData[1].FireSuppressionState = 0
	jsondata.SensorData[1].EventState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
