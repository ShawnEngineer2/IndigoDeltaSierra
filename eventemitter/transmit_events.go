package eventemitter

import (
	"errors"
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sensors"
	"log/slog"
	"strings"

	"indigodeltasierra/customerror"
	"indigodeltasierra/datautil"
)

func TransmitEvents(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, configDS *datamodels.Config, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//This routine manages the transmission of events to the designated output channel

	//Validate the output channel
	switch strings.ToLower(configDS.OutputChannel) {
	case "kafka":
		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Event Output Channel is %s : Broker Endpoint %s : Topic Name \"%s\"", configDS.OutputChannel, configDS.QueueEndpoint, configDS.QueueTopic), false)

	case "filesystem":
		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Event Output Channel is %s : Output Location %s", configDS.OutputChannel, configDS.EventOutputLocation), false)

	default:
		customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Invalid Event Output Channel : %s", configDS.OutputChannel))
		return errors.New("invalid event output channel")
	}

	//Cycle through the Qubz Matrix and transmit the events to the output channel
	for i := 0; i < len(*qubzMatrixCurrent); i++ {

		//Create the Event Header
		eventHeader := datamodels.QubzEventHeader{}

		eventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
		eventHeader.QubzId = (*qubzMatrixCurrent)[i].QubzName
		eventHeader.RouteAssignment = (*qubzMatrixCurrent)[i].RouteAssignment
		eventHeader.ShipmentType = (*qubzMatrixCurrent)[i].ShipmentType
		eventHeader.TransportMode = (*qubzMatrixCurrent)[i].TransportMode

		//Generate and Emit Events
		altimeterEvent := sensors.AltimeterGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(altimeterEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_ALTIMETER, altimeterEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

	}

	return nil

}

func transmitEvent(eventStruct any, configDS *datamodels.Config, qubzName string, sensorType string, eventUUID string, consoleLogger *slog.Logger, fileLogger *slog.Logger) {

	//Transmits the event to the appropriate output channel as indicated by Config
	customlog.CalloutConsole(consoleLogger, (configDS.EventOutputLocation + "/" + eventUUID + ".json"))

	if strings.ToLower(configDS.OutputChannel) == "kafka" {
		customerror.CheckAndPanic(fmt.Errorf("Output Channel \"%s\" not implemented", configDS.OutputChannel))

	} else if strings.ToLower(configDS.OutputChannel) == "filesystem" {
		err := EventToFile(eventStruct, (configDS.EventOutputLocation + "/" + eventUUID))

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Transmission Error : QubzUnit %s / Sensor %s / UUID %s"))
		}
	} else {
		//Invalid output channel - throw error
		customerror.CheckAndPanic(fmt.Errorf("Invalid Output Channel: %s - Cannot Emit Events", configDS.OutputChannel))

	}
}
