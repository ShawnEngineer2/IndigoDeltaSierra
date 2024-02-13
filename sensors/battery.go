package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func BatteryInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Battery sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Battery.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Battery.BatteryChargePct = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_BATTERY_CHARGE_PCT, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Battery.BatterySensorState = 1
	(*qubzMatrix)[matrixIndex].Battery.BatteryTemperature = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_BATTERY_TEMPERATURE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Battery.CapacitanceGelActualVolume = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_CAPACITANCE_GEL_ACTUAL_VOLUME, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.CapacitanceGelExpectedVolume = 136
	(*qubzMatrix)[matrixIndex].Battery.DrainRate = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_DRAIN_RATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.MaxAmpHours = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_MAX_AMP_HOURS, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.RemainingAmpHours = (*qubzMatrix)[matrixIndex].Battery.MaxAmpHours

}

func BatterySet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Battery sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Battery.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Battery.BatteryChargePct = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_BATTERY_CHARGE_PCT, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Battery.BatterySensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_BATTERY_SENSOR_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.BatteryTemperature = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_BATTERY_TEMPERATURE, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Battery.CapacitanceGelActualVolume = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_CAPACITANCE_GEL_ACTUAL_VOLUME, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.CapacitanceGelExpectedVolume = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_CAPACITANCE_GEL_EXPECTED_VOLUME, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.DrainRate = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_DRAIN_RATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.MaxAmpHours = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_BATTERY_MAX_AMP_HOURS, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Battery.RemainingAmpHours = (*qubzMatrix)[matrixIndex].Battery.MaxAmpHours

}
