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

func InfoAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send out INFO style messages to both console and log file
	consoleLogger.Info(msg)
	fileLogger.Info(msg)
}

func ErrorAllChannels(consoleLogger *slog.Logger, fileLogger *slog.Logger, msg string) {

	//Send out ERROR style messages to both console and log file
	consoleLogger.Info(msg)
	fileLogger.Info(msg)
}
