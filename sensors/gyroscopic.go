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

func GyroSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Gyroscopic sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Gyro.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Gyro.GyroscopicSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_GYROSCOPIC_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Gyro.XAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_X_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Gyro.YAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_Y_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Gyro.ZAxisAngle = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_GYROSCOPIC_Z_AXIS_ANGLE, qubzStateDS, consoleLogger, fileLogger)

}

func GyroGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.GyroscopeEvent {

	//Create Event Instance
	eventInstance := datamodels.GyroscopeEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_GYRO
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_GYRO
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.GyroscopeReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Gyro.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].GyroscopicSensorState = (*qubzMatrixCurrent)[matrixIndex].Gyro.GyroscopicSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].XAxisAngle = (*qubzMatrixCurrent)[matrixIndex].Gyro.XAxisAngle
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].YAxisAngle = (*qubzMatrixCurrent)[matrixIndex].Gyro.YAxisAngle
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].ZAxisAngle = (*qubzMatrixCurrent)[matrixIndex].Gyro.ZAxisAngle

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Gyro.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].GyroscopicSensorState = (*qubzMatrixPrevious)[matrixIndex].Gyro.GyroscopicSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].XAxisAngle = (*qubzMatrixPrevious)[matrixIndex].Gyro.XAxisAngle
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].YAxisAngle = (*qubzMatrixPrevious)[matrixIndex].Gyro.YAxisAngle
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].ZAxisAngle = (*qubzMatrixPrevious)[matrixIndex].Gyro.ZAxisAngle

	//Return the completed event
	return eventInstance

}
