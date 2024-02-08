package simulator

import (
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"log/slog"
)

func initializeQubzMatrixExceptions(qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//Set all Qubz in the matrix to a normal state running without exceptions

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Initializing Exception States in Qubz Matrix ...", false)

	for _, x := range *qubzMatrix {
		x.ExceptionAssignment = 0
		x.ExceptionIntermittent = false
		x.ExceptionSeverity = 0
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Exception States set in Qubz Matrix", false)

}
