package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func ComputeTestEvent(outputFilePath string) {
	//This routine is used to test the Compute JSON

	jsondata := datamodels.ComputeEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 3
	jsondata.EventHeader.SensorType.SensorTypeDescription = "compute"

	jsondata.SensorData = make([]datamodels.ComputeReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].AmountMemoryUtilized = 2400
	jsondata.SensorData[0].CPUUtilization = 57
	jsondata.SensorData[0].ComputeFirmwareState = 1
	jsondata.SensorData[0].ComputeFirmwareVersion = 17
	jsondata.SensorData[0].ComputeSensorState = 1
	jsondata.SensorData[0].NumCPUCores = 16
	jsondata.SensorData[0].OSTypeVersion = "Alpine Linux 17"
	jsondata.SensorData[0].RemainingDiskStorage = 2300
	jsondata.SensorData[0].TotalAmountOfMemory = 3600
	jsondata.SensorData[0].TotalDiskStorage = 4000

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].AmountMemoryUtilized = 2400
	jsondata.SensorData[1].CPUUtilization = 57
	jsondata.SensorData[1].ComputeFirmwareState = 1
	jsondata.SensorData[1].ComputeFirmwareVersion = 17
	jsondata.SensorData[1].ComputeSensorState = 1
	jsondata.SensorData[1].NumCPUCores = 16
	jsondata.SensorData[1].OSTypeVersion = "Alpine Linux 17"
	jsondata.SensorData[1].RemainingDiskStorage = 2300
	jsondata.SensorData[1].TotalAmountOfMemory = 3600
	jsondata.SensorData[1].TotalDiskStorage = 4000

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
