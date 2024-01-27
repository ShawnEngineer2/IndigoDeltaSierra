package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func QubzSealTestEvent(outputFilePath string) {
	//This routine is used to test the Qubz Seal JSON

	jsondata := datamodels.SealEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.SealReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].SealEventTime = "2024-01-26T23:38:14.123Z"
	jsondata.SensorData[0].SealSensorState = 1
	jsondata.SensorData[0].SealState = 1

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].SealEventTime = "2024-01-26T23:38:14.123Z"
	jsondata.SensorData[1].SealSensorState = 1
	jsondata.SensorData[1].SealState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
