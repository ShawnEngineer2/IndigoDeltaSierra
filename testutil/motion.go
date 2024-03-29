package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func MotionTestEvent(outputFilePath string) {
	//This routine is used to test the Motion JSON

	jsondata := datamodels.MotionEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 8
	jsondata.EventHeader.SensorType.SensorTypeDescription = "motion sensor"

	jsondata.SensorData = make([]datamodels.MotionReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].AverageVelocity = 2.357
	jsondata.SensorData[0].MotionSensorState = 1
	jsondata.SensorData[0].NumberOfContacts = 3

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].AverageVelocity = 2.357
	jsondata.SensorData[1].MotionSensorState = 1
	jsondata.SensorData[1].NumberOfContacts = 3

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
