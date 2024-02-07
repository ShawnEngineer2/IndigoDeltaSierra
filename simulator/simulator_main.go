package simulator

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"indigodeltasierra/svcclient"
	"indigodeltasierra/sysfile"
	"indigodeltasierra/validators"
	"log/slog"
	"time"
)

func StartSimulation() {
	//This function configures the app and starts the simulation
	//Allocate constants
	const CONFIG_FILE_PATH string = "./config.dat"
	const STARTUP_MSG string = "System Starting"
	const SHUTDOWN_MSG string = "Run Complete ... Shutting Down"
	const DEFAULT_SVC_WAIT int = 10
	var max_qubz_count int = 0

	//Set up console logger
	consoleLogger := slog.Default()
	consoleLogger.Info(STARTUP_MSG)

	//Check for Config file
	consoleLogger.Info("Checking for Config file ... ")

	if !sysfile.FileExists(CONFIG_FILE_PATH) {
		consoleLogger.Error("Cannot find Config file in boot directory ... startup terminated!")
		//return error
	}

	//Load Config file
	consoleLogger.Info("Found Config file ... begin load ...")

	config := datamodels.Config{}

	if !sysfile.LoadFileToStruct(CONFIG_FILE_PATH, &config) {
		consoleLogger.Error("Could not load Config file ... startup terminated")
		//return error

	} else {
		consoleLogger.Info("Config values loaded")
		consoleLogger.Info("Validating Config ...")

		errCount := validators.ValidateConfig(config, *consoleLogger)

		if errCount > 0 {
			consoleLogger.Error(fmt.Sprintf("%s errors identified in loaded config. Startup terminated ... please correct and start again", fmt.Sprint(errCount)))
			//return error
		} else {
			consoleLogger.Info("Config values validated. Loaded values are:")
			consoleLogger.Info(fmt.Sprintf("%+v", config))
		}

	}

	//Setup logfile logger
	consoleLogger.Info("Configuring File Logger to output path " + config.LogLocation + " ...")
	fileLogger := slog.New(slog.NewJSONHandler(customlog.RotatingLog(config.LogLocation), nil))
	fileLogger.Info(STARTUP_MSG)
	fileLogger.Info("Config Values for this run ...")
	fileLogger.Info(fmt.Sprintf("%+v", config))

	//Load Qubz Names and IDs from the Qubz Name File
	qubznames := []datamodels.Qubz{}

	if !sysfile.LoadFileToStruct(config.QubzNameFile, &qubznames) {
		consoleLogger.Error(fmt.Sprintf("Could not load Qubz Names from file %s ... startup terminated", config.QubzNameFile))
		//return error
	} else {
		max_qubz_count = len(qubznames)
		consoleLogger.Info(fmt.Sprintf("Qubz Names loaded from file %s ... %d names loaded", config.QubzNameFile, max_qubz_count))
	}

	//Use the loaded names list to source a random set of Qubz names
	qubzmatrix := make([]datamodels.QubzMatrix, config.QubzCount)

	if config.QubzCount == 0 || config.QubzCount == max_qubz_count {
		//Just load'em all into the Qubz Matrix
		fileLogger.Info("Loading All Qubz Names ...")

		qubzmatrix = make([]datamodels.QubzMatrix, max_qubz_count)

		for i, x := range qubznames {
			qubzmatrix[i].QubzID = x.QubzID
			qubzmatrix[i].QubzName = x.QubzName
		}

		//fmt.Println(qubzmatrix[10])

	} else if config.QubzCount < max_qubz_count {
		//Grab a random set of numbers of QubzCount between 1 and max qubz count and load those names
		fileLogger.Info("Loading Random Qubz Names ...")

		fmt.Println(svcclient.GetRandomNumbers(config.QubzCount, 0, (max_qubz_count - 1), fileLogger, config.EmailAddress))

	} else if config.QubzCount > max_qubz_count {
		//Too many Qubz requested - throw an error
		consoleLogger.Error(fmt.Sprintf("Not enough names in %s file (%d) to satisfy requested number of Qubz (%d) ... startup terminated", config.QubzNameFile, len(qubznames), config.QubzCount))
		//return error
	}

	err := datautil.InitializeQubzMatrix(&qubzmatrix)

	//Cooling off wait so we don't overload the random number generator service
	consoleLogger.Info("Service Cooloff Wait ...")
	time.Sleep(time.Duration(DEFAULT_SVC_WAIT) * time.Second)

	quotaExceeded, err := svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	}

	//Load Location information
	consoleLogger.Info("Loading locations ...")

	locations := make([]datamodels.Location, 1)

	if !sysfile.LoadFileToStruct("./locations.dat", &locations) {
		consoleLogger.Error("Could not load Locations file ... startup terminated")
		//return error
	} else {
		consoleLogger.Info("Location values loaded")
	}

	//Load Route information
	consoleLogger.Info("Loading routes ...")

	routes := make([]datamodels.Route, 1)

	if !sysfile.LoadFileToStruct("./routes.dat", &routes) {
		consoleLogger.Error("Could not load Routes file ... startup terminated")
		//return error
	} else {
		consoleLogger.Info("Route values loaded")
	}

	//Load Sensor Range data
	consoleLogger.Info("Loading sensor ranges ...")

	sensorRanges := make([]datamodels.SensorRange, 1)

	if !sysfile.LoadFileToStruct("./sensor_ranges02.dat", &sensorRanges) {
		consoleLogger.Error("Could not load Sensor Ranges file ... startup terminated")
		//return error
	} else {
		consoleLogger.Info("Sensor Range values loaded")
	}

	fmt.Println(sensorRanges)
	//Randomly assign routes to the Qubz in the Qubz Matrix

	//Cool off wait for random number generator service - then check quota
	consoleLogger.Info("Service Cooloff Wait ...")
	time.Sleep(time.Duration(DEFAULT_SVC_WAIT) * time.Second)

	quotaExceeded, err = svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	}

	//Load the Event Types file into a struct

	quotaExceeded, err = svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		//return error
	}

	//Exit with a CLEAN (no errors) code
	fileLogger.Info(SHUTDOWN_MSG)
	consoleLogger.Info(SHUTDOWN_MSG)
	//return error
}
