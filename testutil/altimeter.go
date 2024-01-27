package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func AltimeterTestEvent(outputFilePath string) {
	//This routine is used to test the altimeter JSON

	altimeterdata := datamodels.AltimeterReading{}

	altimeterdata.AltimeterState = 1
	altimeterdata.Altitude = 2000
	altimeterdata.EventState = 0

	jsondata := datamodels.AltimeterEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.AltimeterReading, 2)

	jsondata.SensorData[0].AltimeterState = 1
	jsondata.SensorData[0].Altitude = 2000
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].AltimeterState = 1
	jsondata.SensorData[1].Altitude = 2330
	jsondata.SensorData[1].EventState = 1

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
