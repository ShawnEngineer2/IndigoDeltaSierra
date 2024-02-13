package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func GyroInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Gyroscopic sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Gyro.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Gyro.GyroscopicSensorState = 1
	(*qubzMatrix)[matrixIndex].Gyro.XAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_X_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Gyro.YAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_Y_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Gyro.ZAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_Z_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)

}
