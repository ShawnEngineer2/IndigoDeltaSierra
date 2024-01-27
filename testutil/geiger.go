package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func GeigerTestEvent(outputFilePath string) {
	//This routine is used to test the Geiger Counter JSON

	jsondata := datamodels.GeigerEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.GeigerReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].GeigerCounterState = 1
	jsondata.SensorData[0].GeigerReading = 435

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].GeigerCounterState = 1
	jsondata.SensorData[1].GeigerReading = 435

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
