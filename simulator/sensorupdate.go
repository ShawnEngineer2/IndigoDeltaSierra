package simulator

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sensors"
	"log/slog"
)

func updateQubzMatrixSensors(qubzMatrix *[]datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//Walk through and set all sensors to initial nominal values - no sensors should start in
	//an exception state

	customlog.InfoConsole(consoleLogger, "Start Sensor Update ....", false)

	var loopCounter int = 0
	var counterResetBoundary int = 1000
	var numQubzProcessed int = 0
	var totalQubzToProcess int = len(*qubzMatrix)

	for i, x := range *qubzMatrix {

		//Create a new sensor state of nominal values
		sensorState, err := CreateSensorState(x, sensorRangeDS, consoleLogger, fileLogger)

		if err != nil {
			return err
		}

		//Use the Sensor State as input into the various sensor update routines
		sensors.AltimeterSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.BatterySet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.ComputeSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.FireSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.GeigerSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		//Skip GPS for now
		sensors.GyroSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.LockSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.MotionSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.QubzSealSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.RadioSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.SpectrometerSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.TempBarometricSet(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)

		//Increment the counters
		loopCounter++
		numQubzProcessed++

		if loopCounter == counterResetBoundary {
			loopCounter = 0
			customlog.InfoConsole(consoleLogger, fmt.Sprintf("%d of %d Qubz Initialized ...", numQubzProcessed, totalQubzToProcess), false)
		}

	}

	customlog.InfoConsole(consoleLogger, "End Sensor Update", false)

	return nil
}
