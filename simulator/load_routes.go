package simulator

import (
	"errors"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sysfile"
	"log/slog"
)

func loadRoutes(targetRoutes *[]datamodels.Route, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {

	//Check for Routes file
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Checking for Routes file "+appconstants.ROUTES_FILE_PATH+" ...")

	if !sysfile.FileExists(appconstants.ROUTES_FILE_PATH) {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Cannot find Routes file in boot directory ... startup terminated!")
		return errors.New("fatal error occurred")
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Loading Routes data ...")

	if !sysfile.LoadFileToStruct(appconstants.ROUTES_FILE_PATH, targetRoutes) {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Could not load Routes file ... startup terminated!")
		return errors.New("fatal error occurred")

	} else {
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Routes values loaded")
	}

	return nil

}
