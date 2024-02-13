package simulator

import (
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"log/slog"
)

func runSimulationCycle(qubzMatrix *[]datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) {

	//This routine runs the simulation cycle

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Assigning Exceptions ...", false)

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Calculating New Sensor State Values ...", false)
	updateQubzMatrixSensors(qubzMatrix, sensorRangeDS, consoleLogger, fileLogger)

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Cycle Complete ... Ready for Event Transmission", false)

}
