package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func QubzSealInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Qubz Seal sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].QubzSeal.SealSensorState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

}

func QubzSealSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Qubz Seal sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].QubzSeal.SealSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SEALSTATE_SEALSENSORSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].QubzSeal.SealState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SEALSTATE_SEALSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

}

func QubzSealGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.SealEvent {

	//Create Event Instance
	eventInstance := datamodels.SealEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_SEAL
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_SEAL
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.SealReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].QubzSeal.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].SealEventTime = (*qubzMatrixCurrent)[matrixIndex].QubzSeal.SealEventTime
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].SealSensorState = (*qubzMatrixCurrent)[matrixIndex].QubzSeal.SealSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].SealState = (*qubzMatrixCurrent)[matrixIndex].QubzSeal.SealState

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].QubzSeal.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].SealEventTime = (*qubzMatrixPrevious)[matrixIndex].QubzSeal.SealEventTime
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].SealSensorState = (*qubzMatrixPrevious)[matrixIndex].QubzSeal.SealSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].SealState = (*qubzMatrixPrevious)[matrixIndex].QubzSeal.SealState

	//Return the completed event
	return eventInstance

}
