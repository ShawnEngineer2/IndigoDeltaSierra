package simulator

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"log/slog"
)

func initializeQubzMatrixExceptions(qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//Set all Qubz in the matrix to a normal state running without exceptions

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Initializing Exception States in Qubz Matrix ...", false)

	for _, x := range *qubzMatrix {
		x.ExceptionAssignment = 0
		x.ExceptionSeverity = appconstants.SENSOR_EXCEPTION_SEVERITY_NONE
		x.ExceptionType = appconstants.SENSOR_EXCEPTION_TYPE_NONE
		x.ExceptionIntervalBoundary = 0
		x.CurrentExceptionInterval = 0
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Exception States set in Qubz Matrix", false)

}
