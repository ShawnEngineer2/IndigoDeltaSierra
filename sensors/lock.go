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

func LockGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.LockEvent {

	//Create Event Instance
	eventInstance := datamodels.LockEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_LOCK
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_LOCK
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.LockReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Lock.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LockEventMethod = (*qubzMatrixCurrent)[matrixIndex].Lock.LockEventMethod
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LockEventTime = (*qubzMatrixCurrent)[matrixIndex].Lock.LockEventTime
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LockEventType = (*qubzMatrixCurrent)[matrixIndex].Lock.LockEventType
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].LockState = (*qubzMatrixCurrent)[matrixIndex].Lock.LockState

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Lock.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LockEventMethod = (*qubzMatrixPrevious)[matrixIndex].Lock.LockEventMethod
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LockEventTime = (*qubzMatrixPrevious)[matrixIndex].Lock.LockEventTime
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LockEventType = (*qubzMatrixPrevious)[matrixIndex].Lock.LockEventType
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].LockState = (*qubzMatrixPrevious)[matrixIndex].Lock.LockState

	//Return the completed event
	return eventInstance

}
