package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func GyroInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Gyroscopic sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Gyro.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Gyro.GyroscopicSensorState = 1
		(*qubzMatrix)[i].Gyro.XAxisAngle = 2.35747
		(*qubzMatrix)[i].Gyro.YAxisAngle = 0.7382
		(*qubzMatrix)[i].Gyro.ZAxisAngle = 0

	}

}
