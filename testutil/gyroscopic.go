package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func GyroscopicTestEvent(outputFilePath string) {
	//This routine is used to test the Gyroscopic JSON

	jsondata := datamodels.GyroscopeEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.GyroscopeReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].XAxisAngle = 3
	jsondata.SensorData[0].YAxisAngle = 2
	jsondata.SensorData[0].ZAxisAngle = 5
	jsondata.SensorData[0].GyroscopicSensorState = 1

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].XAxisAngle = 3
	jsondata.SensorData[1].YAxisAngle = 2
	jsondata.SensorData[1].ZAxisAngle = 5
	jsondata.SensorData[1].GyroscopicSensorState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
