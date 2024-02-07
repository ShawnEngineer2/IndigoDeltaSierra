package customlog

import (
	"io"

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
