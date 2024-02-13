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
