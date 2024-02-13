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
