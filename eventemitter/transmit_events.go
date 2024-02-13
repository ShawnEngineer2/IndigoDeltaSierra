package eventemitter

import (
	"indigodeltasierra/customlog"
	"log/slog"
)

func TransmitEvents(consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Outputting Events ...", false)
}
