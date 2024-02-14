package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func GeigerInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Geiger Counter sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Geiger.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Geiger.GeigerCounterState = 1
	(*qubzMatrix)[matrixIndex].Geiger.GeigerReading = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GEIGER_GEIGERREADING, qubzStateDS, consoleLogger, fileLogger)

}

func GeigerSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Geiger Counter sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Geiger.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Geiger.GeigerCounterState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GEIGER_GEIGERCOUNTERSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Geiger.GeigerReading = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GEIGER_GEIGERREADING, qubzStateDS, consoleLogger, fileLogger)

}

func GeigerGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.GeigerEvent {

	//Create Event Instance
	eventInstance := datamodels.GeigerEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_GEIGER
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_GEIGER
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.GeigerReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Geiger.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].GeigerCounterState = (*qubzMatrixCurrent)[matrixIndex].Geiger.GeigerCounterState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].GeigerReading = (*qubzMatrixCurrent)[matrixIndex].Geiger.GeigerReading

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Geiger.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].GeigerCounterState = (*qubzMatrixPrevious)[matrixIndex].Geiger.GeigerCounterState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].GeigerReading = (*qubzMatrixPrevious)[matrixIndex].Geiger.GeigerReading

	//Return the completed event
	return eventInstance

}
