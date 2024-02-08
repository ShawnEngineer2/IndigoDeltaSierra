package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func TempBarometricTestEvent(outputFilePath string) {
	//This routine is used to test the Temperature & Barometrics JSON

	jsondata := datamodels.TempBarometricsEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 12
	jsondata.EventHeader.SensorType.SensorTypeDescription = "temperature and barometric"

	jsondata.SensorData = make([]datamodels.TempBarometricsReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].BarometricSensorState = 1
	jsondata.SensorData[0].HumidityLevel = 33
	jsondata.SensorData[0].OverallInternalTemperature = 17
	jsondata.SensorData[0].TemperatureSensorState = 1
	jsondata.SensorData[0].MoistureLevel = 1.23479

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].BarometricSensorState = 1
	jsondata.SensorData[1].HumidityLevel = 33
	jsondata.SensorData[1].OverallInternalTemperature = 17
	jsondata.SensorData[1].TemperatureSensorState = 1
	jsondata.SensorData[1].MoistureLevel = 1.23479

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
