package main

import (
	"fmt"
	"indigodeltasierra/SvcClient"
	"io"
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

type config struct {
	intervalSeed    int
	emailAddress    string
	queueEndpointIP string
	queueTopic      string
	logLocation     string
}

func main() {

	//Configure settings for the simulator
	config := config{
		intervalSeed:    15,
		emailAddress:    "shawn.engineer2@gmail.com",
		queueEndpointIP: "",
		queueTopic:      "Topic",
		logLocation:     "output/logs/logfile",
	}

	logger := slog.New(slog.NewJSONHandler(RotatingLog(config.logLocation), nil))

	randomNumberSet := SvcClient.GetRandomNumbers(6, 1, 6, logger, config.emailAddress)

	for _, randomNumber := range randomNumberSet {
		fmt.Println(randomNumber)
	}
}

func RotatingLog(path string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,     // In MB before rotating the file
		MaxAge:     1,     // In days before deleting the file
		MaxBackups: 5,     // Maximum number of backups to keep track of
		Compress:   false, // Compress the rotated log files, false by default.
	}
}
