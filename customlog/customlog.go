package customlog

import (
	"io"
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

func RotatingLog(path string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,     // In MB before rotating the file
		MaxAge:     1,     // In days before deleting the file
		MaxBackups: 5,     // Maximum number of backups to keep track of
		Compress:   false, // Compress the rotated log files, false by default.
	}
}

func InfoAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string, emphasis bool) {

	//Send out an INFO style message to both console and log file
	InfoConsole(consoleLogger, msg, emphasis)
	InfoFile(fileLogger, msg)
}

func InfoConsole(consoleLogger *slog.Logger, msg string, emphasis bool) {

	//Send out INFO style messages to console only
	if emphasis {
		consoleLogger.Info(yellow.Style(msg))
	} else {
		consoleLogger.Info(white.Style(msg))
	}
}

func InfoFile(fileLogger *slog.Logger, msg string) {

	//Send out INFO style messages to log file only
	fileLogger.Info(msg)
}

func ErrorAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send out ERROR style messages to both console and log file
	ErrorConsole(consoleLogger, msg)
	ErrorFile(fileLogger, msg)

}

func ErrorConsole(consoleLogger *slog.Logger, msg string) {

	//Send out ERROR style messages to console only
	consoleLogger.Info(red.Style(msg))

}

func ErrorFile(fileLogger *slog.Logger, msg string) {

	//Send out ERROR style messages to log file only
	fileLogger.Info(msg)

}

func CalloutAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send out an INFO style message to both console and log file
	CalloutConsole(consoleLogger, msg)
	InfoFile(fileLogger, msg)
}

func CalloutConsole(consoleLogger *slog.Logger, msg string) {

	//Send out INFO style messages to console only
	consoleLogger.Info(cyan.Style(msg))
}

func GreenlighAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send out an INFO style message to both console and log file
	GreenlightConsole(consoleLogger, msg)
	InfoFile(fileLogger, msg)
}

func GreenlightConsole(consoleLogger *slog.Logger, msg string) {

	//Send out INFO style messages to console only
	consoleLogger.Info(green.Style(msg))
}
