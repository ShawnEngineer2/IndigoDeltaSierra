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

func BatteryGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.BatteryEvent {

	//Create Event Instance
	eventInstance := datamodels.BatteryEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.TransportMode = eventHeader.TransportMode
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_BATTERY
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_BATTERY
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.BatteryReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].BatteryChargePct = (*qubzMatrixCurrent)[matrixIndex].Battery.BatteryChargePct
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].BatterySensorState = (*qubzMatrixCurrent)[matrixIndex].Battery.BatterySensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Battery.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].BatteryTemperature = (*qubzMatrixCurrent)[matrixIndex].Battery.BatteryTemperature
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].CapacitanceGelActualVolume = (*qubzMatrixCurrent)[matrixIndex].Battery.CapacitanceGelActualVolume
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].CapacitanceGelExpectedVolume = (*qubzMatrixCurrent)[matrixIndex].Battery.CapacitanceGelExpectedVolume
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].DrainRate = (*qubzMatrixCurrent)[matrixIndex].Battery.DrainRate
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].MaxAmpHours = (*qubzMatrixCurrent)[matrixIndex].Battery.MaxAmpHours
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].RemainingAmpHours = (*qubzMatrixCurrent)[matrixIndex].Battery.RemainingAmpHours

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].BatteryChargePct = (*qubzMatrixPrevious)[matrixIndex].Battery.BatteryChargePct
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].BatterySensorState = (*qubzMatrixPrevious)[matrixIndex].Battery.BatterySensorState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Battery.EventState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].BatteryTemperature = (*qubzMatrixPrevious)[matrixIndex].Battery.BatteryTemperature
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].CapacitanceGelActualVolume = (*qubzMatrixPrevious)[matrixIndex].Battery.CapacitanceGelActualVolume
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].CapacitanceGelExpectedVolume = (*qubzMatrixPrevious)[matrixIndex].Battery.CapacitanceGelExpectedVolume
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].DrainRate = (*qubzMatrixPrevious)[matrixIndex].Battery.DrainRate
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].MaxAmpHours = (*qubzMatrixPrevious)[matrixIndex].Battery.MaxAmpHours
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].RemainingAmpHours = (*qubzMatrixPrevious)[matrixIndex].Battery.RemainingAmpHours

	//Return the completed event
	return eventInstance

}
