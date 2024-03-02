package simulator

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sensors"
	"log/slog"
)

//This set of functions initializes the sensors in the Qubz matrix

func initializeQubzMatrixSensors(qubzMatrix *[]datamodels.QubzMatrix, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//Walk through and set all sensors to initial nominal values - no sensors should start in
	//an exception state

	customlog.InfoConsole(consoleLogger, "Start Initialization ....", false)

	var loopCounter int = 0
	var counterResetBoundary int = 1000
	var numQubzProcessed int = 0
	var totalQubzToProcess int = len(*qubzMatrix)

	for i, x := range *qubzMatrix {

		//Create a new sensor state of nominal values
		sensorState, err := createSensorState(x, sensorRangeDS, consoleLogger, fileLogger)

		if err != nil {
			return err
		}

		//Use the Sensor State as input into the various sensor init routines
		//Initialize Altimeter
		sensors.AltimeterInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.BatteryInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.ComputeInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.FireInit(qubzMatrix, i)
		sensors.GeigerInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		//Skip GPS for now
		sensors.GyroInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.LockInit(qubzMatrix, i)
		sensors.MotionInit(qubzMatrix, i)
		sensors.QubzSealInit(qubzMatrix, i)
		sensors.RadioInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.SpectrometerInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)
		sensors.TempBarometricInit(qubzMatrix, i, &sensorState, consoleLogger, fileLogger)

		//Increment the counters
		loopCounter++
		numQubzProcessed++

		if loopCounter == counterResetBoundary {
			loopCounter = 0
			customlog.InfoConsole(consoleLogger, fmt.Sprintf("%d of %d Qubz Initialized ...", numQubzProcessed, totalQubzToProcess), false)
		}

	}

	customlog.InfoConsole(consoleLogger, "End Initialization ....", false)

	return nil
}
