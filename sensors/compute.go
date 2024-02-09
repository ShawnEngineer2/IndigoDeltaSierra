package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func ComputeInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Battery sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Compute.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Compute.AmountMemoryUtilized = 50
		(*qubzMatrix)[i].Compute.CPUUtilization = 35
		(*qubzMatrix)[i].Compute.ComputeFirmwareState = 1
		(*qubzMatrix)[i].Compute.ComputeFirmwareVersion = 347
		(*qubzMatrix)[i].Compute.ComputeSensorState = 1
		(*qubzMatrix)[i].Compute.MemoryUtilization = 32
		(*qubzMatrix)[i].Compute.NumCPUCores = 8
		(*qubzMatrix)[i].Compute.OSTypeVersion = "Raspberry Pi 3.47 / GPU Gennex 18.9.17 / Alpine Linux 17.3"
		(*qubzMatrix)[i].Compute.RemainingDiskStorage = 380
		(*qubzMatrix)[i].Compute.TotalAmountOfMemory = 1000
		(*qubzMatrix)[i].Compute.TotalDiskStorage = 500

	}

}
