package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func MotionInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Motion sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Motion.MotionSensorState = 1
	(*qubzMatrix)[matrixIndex].Motion.NumberOfContacts = 0
	(*qubzMatrix)[matrixIndex].Motion.AverageVelocity = 0

}

func MotionSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Motion sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Motion.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Motion.MotionSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_MOTIONSENSORSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Motion.NumberOfContacts = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_NUMBEROFCONTACTS, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Motion.AverageVelocity = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_MOTION_AVERAGEVELOCITY, qubzStateDS, consoleLogger, fileLogger)

}

func MotionGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.MotionEvent {

	//Create Event Instance
	eventInstance := datamodels.MotionEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_MOTION
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_MOTION
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.MotionReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].AverageVelocity = (*qubzMatrixCurrent)[matrixIndex].Motion.AverageVelocity
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Motion.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].MotionSensorState = (*qubzMatrixCurrent)[matrixIndex].Motion.MotionSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].NumberOfContacts = (*qubzMatrixCurrent)[matrixIndex].Motion.NumberOfContacts

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].AverageVelocity = (*qubzMatrixPrevious)[matrixIndex].Motion.AverageVelocity
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Motion.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].MotionSensorState = (*qubzMatrixPrevious)[matrixIndex].Motion.MotionSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].NumberOfContacts = (*qubzMatrixPrevious)[matrixIndex].Motion.NumberOfContacts

	//Return the completed event
	return eventInstance

}
