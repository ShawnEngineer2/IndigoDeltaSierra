package eventemitter

import (
	"errors"
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/sensors"
	"indigodeltasierra/svcclient"
	"indigodeltasierra/svcclient/models"
	"log/slog"
	"strings"

	"encoding/json"
	"indigodeltasierra/customerror"
	"indigodeltasierra/datautil"
)

func TransmitEvents(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, configDS *datamodels.Config, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//This routine manages the transmission of events to the designated output channel

	//Validate the output channel
	switch strings.ToLower(configDS.OutputChannel) {
	case "segmentio":
		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Event Output Channel is %s : Broker Endpoint %s : Topic Name \"%s\"", configDS.OutputChannel, configDS.QueueEndpoint, configDS.QueueTopic), false)

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

		batteryEvent := sensors.BatteryGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(batteryEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_BATTERY, batteryEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		computeEvent := sensors.ComputeGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(computeEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_COMPUTE, computeEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		fireEvent := sensors.FireGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(fireEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_FIRE, fireEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		geigerEvent := sensors.GeigerGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(geigerEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_GEIGER, geigerEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		gyroEvent := sensors.GyroGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(gyroEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_GYRO, gyroEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		lockEvent := sensors.LockGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(lockEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_LOCK, lockEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		qubzSealEvent := sensors.QubzSealGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(qubzSealEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_SEAL, qubzSealEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		radioEvent := sensors.RadioGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(radioEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_RADIO, radioEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		spectrometerEvent := sensors.SpectrometerGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(spectrometerEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_SPECTROMETER, spectrometerEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		tempBarometricEvent := sensors.TempBarometricGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(tempBarometricEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_TEMPHUMIDITY, tempBarometricEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

		motionEvent := sensors.MotionGetEvent(qubzMatrixCurrent, qubzMatrixPrevious, i, &eventHeader, consoleLogger, fileLogger)
		transmitEvent(motionEvent, configDS, eventHeader.QubzId, appconstants.SENSOR_TYPE_MOTION, motionEvent.EventHeader.EventUUID, consoleLogger, fileLogger)

	}

	return nil

}

func transmitEvent(eventStruct any, configDS *datamodels.Config, qubzName string, sensorType string, eventUUID string, consoleLogger *slog.Logger, fileLogger *slog.Logger) {

	//Transmits the event to the appropriate output channel as indicated by Config
	if strings.ToLower(configDS.OutputChannel) == "segmentio" {

		//Convert the event struct to a JSON string
		jsonbytes, err := json.Marshal(eventStruct)

		if err != nil {
			fmt.Println(err.Error())
		}

		//Use the passed Sensor Type string and the Sensor ID to determine the target partition
		sensorID := datautil.GetSensorIDForSensorDesc(sensorType)
		partitionID := sensorID - 1

		if partitionID < 0 || partitionID > 12 {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Transmission Error : Could not identify target partition for Sensor Type %s. Data Not Transmitted for this sensor.", sensorType))
			return
		}

		//Create an event structure
		eventData := models.EventData{
			EventKey:        sensorType,
			EventData:       string(jsonbytes),
			TargetPartition: partitionID,
		}

		//Send the message
		svcclient.S_produceKafkaMessage(&eventData, configDS.QueueTopic, configDS.QueueEndpoint)

	} else if strings.ToLower(configDS.OutputChannel) == "kafka" {

		//Convert the event struct to a JSON string
		jsonbytes, err := json.Marshal(eventStruct)

		if err != nil {
			fmt.Println(err.Error())
		}

		//Use the passed Sensor Type string and the Sensor ID to determine the target partition
		sensorID := datautil.GetSensorIDForSensorDesc(sensorType)
		partitionID := sensorID - 1

		if partitionID < 0 || partitionID > 12 {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Transmission Error : Could not identify target partition for Sensor Type %s. Data Not Transmitted for this sensor.", sensorType))
			return
		}

		//Create an event structure
		eventData := models.EventData{
			EventKey:        sensorType,
			EventData:       string(jsonbytes),
			TargetPartition: partitionID,
		}

		//Send the message
		svcclient.ProduceSingleKafkaMessage(&eventData, configDS.QueueTopic, configDS.QueueEndpoint)

	} else if strings.ToLower(configDS.OutputChannel) == "filesystem" {
		err := EventToFile(eventStruct, (configDS.EventOutputLocation + "/" + eventUUID))

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Transmission Error : QubzUnit \"%s\" / Sensor \"%s\" / UUID \"%s\"", qubzName, sensorType, eventUUID))
		}
	} else {
		//Invalid output channel - throw error
		customerror.CheckAndPanic(fmt.Errorf("invalid Output Channel: %s - Cannot Emit Events", configDS.OutputChannel))

	}
}
