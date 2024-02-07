package datautil

import (
	"errors"
	"indigodeltasierra/customlog"
	"indigodeltasierra/sysfile"
	"log/slog"
)

// Message routing constants
const route_TO_INFO int = 0
const route_TO_ERROR int = 1

func LoadDataFile(targetStruct any, dataDesc string, filePath string, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {

	//Verify data description and use default if needed
	if len(dataDesc) == 0 {
		dataDesc = "data"
	}

	//Check for data file
	doMessage(route_TO_INFO, consoleLogger, fileLogger, "Checking for "+dataDesc+" file "+filePath+" ...")

	if !sysfile.FileExists(filePath) {
		doMessage(route_TO_ERROR, consoleLogger, fileLogger, "Cannot find "+dataDesc+" file in boot directory ... startup terminated!")
		return errors.New("fatal error occurred")
	}

	doMessage(route_TO_INFO, consoleLogger, fileLogger, "Loading "+dataDesc+" file ...")

	if !sysfile.LoadFileToStruct(filePath, targetStruct) {
		doMessage(route_TO_ERROR, consoleLogger, fileLogger, "Could not load "+dataDesc+" file ... startup terminated!")
		return errors.New("fatal error occurred")

	} else {
		doMessage(route_TO_INFO, consoleLogger, fileLogger, dataDesc+" file loaded")
	}

	return nil

}

func doMessage(routing int, consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send message to user
	if routing == route_TO_ERROR {
		doErrorMessage(consoleLogger, fileLogger, msg)

	} else {
		doInfoMessage(consoleLogger, fileLogger, msg)

	}

}

func doInfoMessage(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//This routine simplifies handling of Info messages based on if a File Logger is present or not
	if fileLogger == nil {

		customlog.InfoConsole(consoleLogger, msg, false)

	} else {

		customlog.InfoAllChannels(consoleLogger, fileLogger, msg, false)

	}

}

func doErrorMessage(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//This routine simplifies handling of Error messages based on if a File Logger is present or not
	if fileLogger == nil {

		customlog.ErrorConsole(consoleLogger, msg)

	} else {

		customlog.ErrorAllChannels(consoleLogger, fileLogger, msg)

	}

}
