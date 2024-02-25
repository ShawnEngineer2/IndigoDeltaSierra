package testutil

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func CreateTestException() {

	//Create some test exceptions
	exceptionList := make([]datamodels.QubzException, 3)

	exceptionList[0].ExceptionId = 0
	exceptionList[0].ExceptionDesc = "Test Exception 1"
	exceptionList[0].SeverityLevel = 3
	exceptionList[0].ExceptionType = appconstants.SENSOR_EXCEPTION_TYPE_CONTINUOUS
	exceptionList[0].IntermittencyInterval = 0

	exceptionList[0].AffectedSensors = make([]datamodels.QubzSensorDataPointException, 3)

	exceptionList[0].AffectedSensors[0].SensorDataPointId = 143
	exceptionList[0].AffectedSensors[0].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[0].AffectedSensors[0].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[0].AffectedSensors[0].FixedModValue = 0

	exceptionList[0].AffectedSensors[1].SensorDataPointId = 143
	exceptionList[0].AffectedSensors[1].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[0].AffectedSensors[1].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[0].AffectedSensors[1].FixedModValue = 0

	exceptionList[0].AffectedSensors[2].SensorDataPointId = 143
	exceptionList[0].AffectedSensors[2].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[0].AffectedSensors[2].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[0].AffectedSensors[2].FixedModValue = 0

	exceptionList[1].ExceptionId = 1
	exceptionList[1].ExceptionDesc = "Test Exception 2"
	exceptionList[1].SeverityLevel = 3
	exceptionList[1].ExceptionType = appconstants.SENSOR_EXCEPTION_TYPE_CONTINUOUS
	exceptionList[1].IntermittencyInterval = 0

	exceptionList[1].AffectedSensors = make([]datamodels.QubzSensorDataPointException, 3)

	exceptionList[1].AffectedSensors[0].SensorDataPointId = 143
	exceptionList[1].AffectedSensors[0].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[1].AffectedSensors[0].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[1].AffectedSensors[0].FixedModValue = 0

	exceptionList[1].AffectedSensors[1].SensorDataPointId = 143
	exceptionList[1].AffectedSensors[1].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[1].AffectedSensors[1].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[1].AffectedSensors[1].FixedModValue = 0

	exceptionList[1].AffectedSensors[2].SensorDataPointId = 143
	exceptionList[1].AffectedSensors[2].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[1].AffectedSensors[2].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[1].AffectedSensors[2].FixedModValue = 0

	exceptionList[2].ExceptionId = 2
	exceptionList[2].ExceptionDesc = "Test Exception 3"
	exceptionList[2].SeverityLevel = 3
	exceptionList[2].ExceptionType = appconstants.SENSOR_EXCEPTION_TYPE_CONTINUOUS
	exceptionList[2].IntermittencyInterval = 0

	exceptionList[2].AffectedSensors = make([]datamodels.QubzSensorDataPointException, 3)

	exceptionList[2].AffectedSensors[0].SensorDataPointId = 143
	exceptionList[2].AffectedSensors[0].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[2].AffectedSensors[0].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[2].AffectedSensors[0].FixedModValue = 0

	exceptionList[2].AffectedSensors[1].SensorDataPointId = 143
	exceptionList[2].AffectedSensors[1].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[2].AffectedSensors[1].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[2].AffectedSensors[1].FixedModValue = 0

	exceptionList[2].AffectedSensors[2].SensorDataPointId = 143
	exceptionList[2].AffectedSensors[2].SensorDataPointDesc = "Fire Suppression State"
	exceptionList[2].AffectedSensors[2].ValueModType = appconstants.SENSOR_EXCEPTION_VALUE_MOD_LOW
	exceptionList[2].AffectedSensors[2].FixedModValue = 0

	err := eventemitter.EventToFile(exceptionList, "./output/events/exception.json")

	customerror.CheckAndPanic(err)

}
