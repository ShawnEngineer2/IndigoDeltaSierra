package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func LockTestEvent(outputFilePath string) {
	//This routine is used to test the Lock JSON

	jsondata := datamodels.LockEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 7
	jsondata.EventHeader.SensorType.SensorTypeDescription = "lock state"

	jsondata.SensorData = make([]datamodels.LockReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].LockEventMethod = 1
	jsondata.SensorData[0].LockEventTime = "2024-01-26T23:38:14.123Z"
	jsondata.SensorData[0].LockEventType = 1
	jsondata.SensorData[0].LockState = 1

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].LockEventMethod = 1
	jsondata.SensorData[1].LockEventTime = "2024-01-26T23:38:14.123Z"
	jsondata.SensorData[1].LockEventType = 1
	jsondata.SensorData[1].LockState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
