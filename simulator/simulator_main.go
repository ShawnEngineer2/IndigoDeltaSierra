package simulator

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"log/slog"
)

func StartSimulation() {
	//This function configures the app and starts the simulation

	//Declare data structures to be used
	configDS := datamodels.Config{}
	locationsDS := make([]datamodels.Location, 1)
	routesDS := make([]datamodels.Route, 1)

	//Set up console logger
	consoleLogger := slog.Default()
	consoleLogger.Info(appconstants.STARTUP_MSG)

	//Load the Config file
	err := loadConfig(&configDS, consoleLogger)

	if err != nil {
		return
	}

	//Setup logfile logger
	consoleLogger.Info("Configuring File Logger to output path " + configDS.LogLocation + " ...")
	fileLogger := slog.New(slog.NewJSONHandler(customlog.RotatingLog(configDS.LogLocation), nil))
	fileLogger.Info(appconstants.STARTUP_MSG)
	fileLogger.Info("Config Values for this run ...")
	fileLogger.Info(fmt.Sprintf("%+v", configDS))

	//Load Locations data
	err = loadLocations(&locationsDS, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load Route data

	err = loadRoutes(&routesDS, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	/*
		//Cooling off wait so we don't overload the random number generator service
		consoleLogger.Info("Service Cooloff Wait ...")
		time.Sleep(time.Duration(appconstants.DEFAULT_SVC_WAIT) * time.Second)

		quotaExceeded, err := svcclient.CheckQuotaExceeded(fileLogger, configDS.EmailAddress)

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
		time.Sleep(time.Duration(appconstants.DEFAULT_SVC_WAIT) * time.Second)

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
	*/
	//Exit with a CLEAN (no errors) code
	customlog.InfoAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_COMPLETE_MSG)

}
