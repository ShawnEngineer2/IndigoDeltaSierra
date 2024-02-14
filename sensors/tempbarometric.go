package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func TempBarometricInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Temperature and Barometric sensors

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].TempBarometrics.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].TempBarometrics.BarometricSensorState = 1
	(*qubzMatrix)[matrixIndex].TempBarometrics.HumidityLevel = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_TEMPANDHUMIDITY_HUMIDITY_LEVEL, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].TempBarometrics.MoistureLevel = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_TEMPANDHUMIDITY_MOISTURE_LEVEL, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].TempBarometrics.OverallInternalTemperature = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_TEMPANDHUMIDITY_OVERALL_INTERNAL_TEMPERATURE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].TempBarometrics.TemperatureSensorState = 1

}

func TempBarometricSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {

	//Call the Init routine to set most values
	TempBarometricInit(qubzMatrix, matrixIndex, qubzStateDS, consoleLogger, fileLogger)

	//Set the data points not treated as variable-value by INIT
	(*qubzMatrix)[matrixIndex].TempBarometrics.BarometricSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_TEMPANDHUMIDITY_BAROMETRIC_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].TempBarometrics.TemperatureSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_TEMPANDHUMIDITY_TEMPERATURE_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))

}

func TempBarometricGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.TempBarometricsEvent {

	//Create Event Instance
	eventInstance := datamodels.TempBarometricsEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_TEMPHUMIDITY
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_TEMPHUMIDITY
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.TempBarometricsReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].BarometricSensorState = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.BarometricSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].HumidityLevel = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.HumidityLevel
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].MoistureLevel = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.MoistureLevel
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].OverallInternalTemperature = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.OverallInternalTemperature
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].TemperatureSensorState = (*qubzMatrixCurrent)[matrixIndex].TempBarometrics.TemperatureSensorState

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].BarometricSensorState = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.BarometricSensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].HumidityLevel = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.HumidityLevel
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].MoistureLevel = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.MoistureLevel
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].OverallInternalTemperature = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.OverallInternalTemperature
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].TemperatureSensorState = (*qubzMatrixPrevious)[matrixIndex].TempBarometrics.TemperatureSensorState

	//Return the completed event
	return eventInstance

}
