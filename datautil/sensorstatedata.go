package datautil

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"log/slog"
)

func GetSensorStateValue(sensorDataPointId int, sensorState *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) float64 {
	//Find the passed datapoint ID in the referenced Sensor State

	for _, x := range (*sensorState).SensorDataPoints {

		if x.SensorDataPointId == sensorDataPointId {
			return x.DataPointValue
		}

	}

	//If you get this far you didn't find anything - note it in the log
	customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Cannot find Data Point ID %d in Sensor State buffer!", sensorDataPointId))
	return 0
}
