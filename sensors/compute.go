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
