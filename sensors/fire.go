package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func FireInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Fire sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Fire.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Fire.FireDetected = 0
	(*qubzMatrix)[matrixIndex].Fire.FireSensorState = 1
	(*qubzMatrix)[matrixIndex].Fire.FireSuppressionState = 0

}

func FireSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Fire sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Fire.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Fire.FireDetected = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_FIRE_FIREDETECTED, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Fire.FireSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_FIRE_FIRESENSORSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Fire.FireSuppressionState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_FIRE_FIRESUPPRESSIONSTATE, qubzStateDS, consoleLogger, fileLogger))

}

func FireGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.FireEvent {

	//Create Event Instance
	eventInstance := datamodels.FireEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_FIRE
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_FIRE
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.FireReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].FireDetected = (*qubzMatrixCurrent)[matrixIndex].Fire.FireDetected
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].FireSensorState = (*qubzMatrixCurrent)[matrixIndex].Fire.FireSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].FireSuppressionState = (*qubzMatrixCurrent)[matrixIndex].Fire.FireSuppressionState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Fire.EventState

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].FireDetected = (*qubzMatrixPrevious)[matrixIndex].Fire.FireDetected
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].FireSensorState = (*qubzMatrixPrevious)[matrixIndex].Fire.FireSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].FireSuppressionState = (*qubzMatrixPrevious)[matrixIndex].Fire.FireSuppressionState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Fire.EventState

	//Return the completed event
	return eventInstance

}
