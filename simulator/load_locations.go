package simulator

import (
	"errors"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sysfile"
	"log/slog"
)

func loadLocations(targetLocations *[]datamodels.Location, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {

	//Check for Locations file
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Checking for Locations file "+appconstants.LOCATIONS_FILE_PATH+" ...")

	if !sysfile.FileExists(appconstants.LOCATIONS_FILE_PATH) {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Cannot find Locations file in boot directory ... startup terminated!")
		return errors.New("fatal error occurred")
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Loading Locations data ...")

	if !sysfile.LoadFileToStruct(appconstants.LOCATIONS_FILE_PATH, targetLocations) {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Could not load Locations file ... startup terminated!")
		return errors.New("fatal error occurred")

	} else {
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Location values loaded")
	}

	return nil

}
