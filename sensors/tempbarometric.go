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
