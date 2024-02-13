package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func LockInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Lock sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Lock.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Lock.LockEventTime = "02/08/2024 21:37:10.093"
	(*qubzMatrix)[matrixIndex].Lock.LockEventMethod = 3
	(*qubzMatrix)[matrixIndex].Lock.LockEventType = 1
	(*qubzMatrix)[matrixIndex].Lock.LockState = 1

}

func LockSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Lock sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Lock.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Lock.LockEventTime = "02/08/2024 21:37:10.093"
	(*qubzMatrix)[matrixIndex].Lock.LockEventMethod = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_LOCK_LOCK_EVENT_METHOD, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Lock.LockEventType = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_LOCK_LOCK_EVENT_TYPE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Lock.LockState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_LOCK_LOCK_STATE, qubzStateDS, consoleLogger, fileLogger))

}
