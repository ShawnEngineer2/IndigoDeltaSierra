package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func BatteryTestEvent(outputFilePath string) {
	//This routine is used to test the battery JSON

	jsondata := datamodels.BatteryEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 2
	jsondata.EventHeader.SensorType.SensorTypeDescription = "battery"

	jsondata.SensorData = make([]datamodels.BatteryReading, 2)

	jsondata.SensorData[0].BatterySensorState = 1
	jsondata.SensorData[0].BatteryTemperature = 15
	jsondata.SensorData[0].CapacitanceGelActualVolume = 300
	jsondata.SensorData[0].CapacitanceGelExpectedVolume = 300
	jsondata.SensorData[0].DrainRate = 134
	jsondata.SensorData[0].MaxAmpHours = 50000
	jsondata.SensorData[0].RemainingAmpHours = 34000
	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].BatteryChargePct = 45.78219

	jsondata.SensorData[1].BatterySensorState = 1
	jsondata.SensorData[1].BatteryTemperature = 15
	jsondata.SensorData[1].CapacitanceGelActualVolume = 300
	jsondata.SensorData[1].CapacitanceGelExpectedVolume = 300
	jsondata.SensorData[1].DrainRate = 134
	jsondata.SensorData[1].MaxAmpHours = 50000
	jsondata.SensorData[1].RemainingAmpHours = 34000
	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].BatteryChargePct = 45.78219

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
