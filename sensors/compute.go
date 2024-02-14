package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func ComputeInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Battery sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Compute.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Compute.AmountMemoryUtilized = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_AMOUNT_MEMORY_USED, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Compute.CPUUtilization = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_CPU_UTILIZATION, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Compute.ComputeFirmwareState = 1
	(*qubzMatrix)[matrixIndex].Compute.ComputeFirmwareVersion = 347
	(*qubzMatrix)[matrixIndex].Compute.ComputeSensorState = 1
	(*qubzMatrix)[matrixIndex].Compute.TotalAmountOfMemory = 1000
	(*qubzMatrix)[matrixIndex].Compute.MemoryUtilization = float64((*qubzMatrix)[matrixIndex].Compute.AmountMemoryUtilized) / float64((*qubzMatrix)[matrixIndex].Compute.TotalAmountOfMemory)
	(*qubzMatrix)[matrixIndex].Compute.NumCPUCores = 8
	(*qubzMatrix)[matrixIndex].Compute.OSTypeVersion = "Raspberry Pi 3.47 / GPU Gennex 18.9.17 / Alpine Linux 17.3"
	(*qubzMatrix)[matrixIndex].Compute.RemainingDiskStorage = 380
	(*qubzMatrix)[matrixIndex].Compute.TotalDiskStorage = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_TOTAL_DISK_STORAGE, qubzStateDS, consoleLogger, fileLogger))

}

func ComputeSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Battery sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Compute.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Compute.AmountMemoryUtilized = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_AMOUNT_MEMORY_USED, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Compute.CPUUtilization = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_CPU_UTILIZATION, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Compute.ComputeFirmwareState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_COMPUTE_FIRMWARE_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Compute.ComputeFirmwareVersion = 347
	(*qubzMatrix)[matrixIndex].Compute.ComputeSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_COMPUTE_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Compute.TotalAmountOfMemory = 1000
	(*qubzMatrix)[matrixIndex].Compute.MemoryUtilization = float64((*qubzMatrix)[matrixIndex].Compute.AmountMemoryUtilized) / float64((*qubzMatrix)[matrixIndex].Compute.TotalAmountOfMemory)
	(*qubzMatrix)[matrixIndex].Compute.NumCPUCores = 8
	(*qubzMatrix)[matrixIndex].Compute.OSTypeVersion = "Raspberry Pi 3.47 / GPU Gennex 18.9.17 / Alpine Linux 17.3"
	(*qubzMatrix)[matrixIndex].Compute.RemainingDiskStorage = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_REMAINING_DISK_STORAGE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Compute.TotalDiskStorage = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_COMPUTE_TOTAL_DISK_STORAGE, qubzStateDS, consoleLogger, fileLogger))

}

func ComputeGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.ComputeEvent {

	//Create Event Instance
	eventInstance := datamodels.ComputeEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_COMPUTE
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_COMPUTE
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.ComputeReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].AmountMemoryUtilized = (*qubzMatrixCurrent)[matrixIndex].Compute.AmountMemoryUtilized
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].CPUUtilization = (*qubzMatrixCurrent)[matrixIndex].Compute.CPUUtilization
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Compute.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].ComputeFirmwareState = (*qubzMatrixCurrent)[matrixIndex].Compute.ComputeFirmwareState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].ComputeFirmwareVersion = (*qubzMatrixCurrent)[matrixIndex].Compute.ComputeFirmwareVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].ComputeSensorState = (*qubzMatrixCurrent)[matrixIndex].Compute.ComputeSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].MemoryUtilization = (*qubzMatrixCurrent)[matrixIndex].Compute.MemoryUtilization
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].NumCPUCores = (*qubzMatrixCurrent)[matrixIndex].Compute.NumCPUCores
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].OSTypeVersion = (*qubzMatrixCurrent)[matrixIndex].Compute.OSTypeVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RemainingDiskStorage = (*qubzMatrixCurrent)[matrixIndex].Compute.RemainingDiskStorage
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].TotalAmountOfMemory = (*qubzMatrixCurrent)[matrixIndex].Compute.TotalAmountOfMemory
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].TotalDiskStorage = (*qubzMatrixCurrent)[matrixIndex].Compute.TotalDiskStorage

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].AmountMemoryUtilized = (*qubzMatrixPrevious)[matrixIndex].Compute.AmountMemoryUtilized
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].CPUUtilization = (*qubzMatrixPrevious)[matrixIndex].Compute.CPUUtilization
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Compute.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].ComputeFirmwareState = (*qubzMatrixPrevious)[matrixIndex].Compute.ComputeFirmwareState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].ComputeFirmwareVersion = (*qubzMatrixPrevious)[matrixIndex].Compute.ComputeFirmwareVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].ComputeSensorState = (*qubzMatrixPrevious)[matrixIndex].Compute.ComputeSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].MemoryUtilization = (*qubzMatrixPrevious)[matrixIndex].Compute.MemoryUtilization
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].NumCPUCores = (*qubzMatrixPrevious)[matrixIndex].Compute.NumCPUCores
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].OSTypeVersion = (*qubzMatrixPrevious)[matrixIndex].Compute.OSTypeVersion
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RemainingDiskStorage = (*qubzMatrixPrevious)[matrixIndex].Compute.RemainingDiskStorage
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].TotalAmountOfMemory = (*qubzMatrixPrevious)[matrixIndex].Compute.TotalAmountOfMemory
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].TotalDiskStorage = (*qubzMatrixPrevious)[matrixIndex].Compute.TotalDiskStorage

	//Return the completed event
	return eventInstance

}
