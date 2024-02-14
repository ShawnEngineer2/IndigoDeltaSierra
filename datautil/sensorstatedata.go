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

func AssignEventStatesAll(qubzMatrix *[]datamodels.QubzMatrix, newEventState int) {
	//Change event state in the passed qubzMatrix to the passed new Event State value

	for i := 0; i < len(*qubzMatrix); i++ {

		(*qubzMatrix)[i].Altimeter.EventState = newEventState
		(*qubzMatrix)[i].Battery.EventState = newEventState
		(*qubzMatrix)[i].Compute.EventState = newEventState
		(*qubzMatrix)[i].Fire.EventState = newEventState
		(*qubzMatrix)[i].GPS.EventState = newEventState
		(*qubzMatrix)[i].Geiger.EventState = newEventState
		(*qubzMatrix)[i].Gyro.EventState = newEventState
		(*qubzMatrix)[i].Lock.EventState = newEventState
		(*qubzMatrix)[i].Motion.EventState = newEventState
		(*qubzMatrix)[i].QubzSeal.EventState = newEventState
		(*qubzMatrix)[i].Radio.EventState = newEventState
		(*qubzMatrix)[i].Spectrometer.EventState = newEventState
		(*qubzMatrix)[i].TempBarometrics.EventState = newEventState
	}
}
