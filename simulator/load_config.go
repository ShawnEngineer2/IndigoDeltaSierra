package simulator

import (
	"errors"
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sysfile"
	"indigodeltasierra/validators"
	"log/slog"
)

func loadConfig(targetConfig *datamodels.Config, consoleLogger *slog.Logger) error {
	//Load the configuration file into the passed struct

	//Check for Config file
	consoleLogger.Info("Checking for Config file ... ")

	if !sysfile.FileExists(appconstants.CONFIG_FILE_PATH) {
		consoleLogger.Error("Cannot find Config file in boot directory ... startup terminated!")
		return errors.New("fatal error occurred")
	}

	//Load Config file
	consoleLogger.Info("Found Config file ... begin load ...")

	if !sysfile.LoadFileToStruct(appconstants.CONFIG_FILE_PATH, targetConfig) {
		consoleLogger.Error("Could not load Config file ... startup terminated")
		return errors.New("fatal error occurred")

	} else {
		consoleLogger.Info("Config values loaded")
		consoleLogger.Info("Validating Config ...")

		errCount := validators.ValidateConfig(*targetConfig, *consoleLogger)

		if errCount > 0 {
			consoleLogger.Error(fmt.Sprintf("%s errors identified in loaded config. Startup terminated ... please correct and start again", fmt.Sprint(errCount)))
			return errors.New("fatal error occurred")

		} else {
			consoleLogger.Info("Config values validated. Loaded values are:")
			consoleLogger.Info(fmt.Sprintf("%+v", *targetConfig))
		}

	}

	//All good - return
	return nil

}
