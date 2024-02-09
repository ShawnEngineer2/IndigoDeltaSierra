package simulator

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sensors"
	"log/slog"
)

const init_MSG string = "Initializing Sensor : "
const init_complete_MSG string = "Initialization Complete"

//This set of functions initializes the sensors in the Qubz matrix

func initializeQubzMatrixSensors(qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//Walk through and set all sensors to initial nominal values - no sensors should start in
	//an exception state

	//Note: You'll see that we're making multiple passes through the same dataset which is highly inefficient. This is a DELIBERATE
	//action for a very impractical reason - I want to be able to show messages like "Initializing Sensor: Altimeter" so the user
	//can see progress and it's easier for a developer to identify where something blew up. Since this whole initialization process
	//only occurs once I see it as acceptable so as to satisfy my vanity LOL

	//Initialize Altimeter
	startMessage("Altimeter", consoleLogger, fileLogger)
	sensors.AltimeterInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Battery Sensor
	startMessage("Battery", consoleLogger, fileLogger)
	sensors.BatteryInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Compute Sensors
	startMessage("Onboard Compute", consoleLogger, fileLogger)
	sensors.ComputeInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Fire Sensors
	startMessage("Fire", consoleLogger, fileLogger)
	sensors.FireInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Geiger Counter Sensors
	startMessage("Geiger", consoleLogger, fileLogger)
	sensors.GeigerInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize GPS Counter Sensors
	startMessage("GPS", consoleLogger, fileLogger)
	sensors.GPSInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Gyroscopic Sensors
	startMessage("Gyroscopic", consoleLogger, fileLogger)
	sensors.GyroInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Lock Sensors
	startMessage("Lock", consoleLogger, fileLogger)
	sensors.LockInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Motion Sensors
	startMessage("Motion", consoleLogger, fileLogger)
	sensors.MotionInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Qubz Seal Detectors
	startMessage("Qubz Seal State", consoleLogger, fileLogger)
	sensors.QubzSealInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Radio Detectors
	startMessage("Radio", consoleLogger, fileLogger)
	sensors.RadioInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

	//Initialize Temperature and Barometric Sensors
	startMessage("Temperature and Barometric", consoleLogger, fileLogger)
	sensors.TempBarometricInit(qubzMatrix)
	endMessage(consoleLogger, fileLogger)

}

func startMessage(sensorName string, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("%s %s ...", init_MSG, sensorName))
}

func endMessage(consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	customlog.InfoAllChannels(consoleLogger, fileLogger, init_complete_MSG, false)
}
