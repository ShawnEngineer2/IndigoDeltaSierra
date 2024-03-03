package simulator

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"indigodeltasierra/eventemitter"
	"indigodeltasierra/svcclient"
	clientmodels "indigodeltasierra/svcclient/models"
	"log/slog"
	"strings"
	"time"
)

func StartSimulation() {
	//This function configures the app and starts the simulation

	//Declare data structures to be used
	configDS := datamodels.Config{}
	locationsDS := make([]datamodels.Location, 1)
	routesDS := make([]datamodels.Route, 1)
	classOfServiceDS := make([]datamodels.ClassOfService, 1)
	qubzNameDS := make([]datamodels.Qubz, 1)
	sensorTypeDS := make([]datamodels.SensorType, 1)
	shipmentTypeDS := make([]datamodels.ShipmentType, 1)
	transportModeDS := make([]datamodels.TransportMode, 1)
	sensorRangeDS := make([]datamodels.SensorRange, 1)
	sensorExceptionDS := make([]datamodels.QubzException, 1)
	skafkaConnections := make([]clientmodels.S_kafkaConnection, 13)

	//Set up console logger
	consoleLogger := slog.Default()
	customlog.InfoConsole(consoleLogger, appconstants.STARTUP_MSG, true)

	//Load the Config file
	err := datautil.LoadDataFile(&configDS, "Configuration", appconstants.CONFIG_FILE_PATH, consoleLogger, nil)

	if err != nil {
		return
	}

	//Setup logfile logger
	customlog.InfoConsole(consoleLogger, fmt.Sprintf("Configuring File Logger to output path %s ...", configDS.LogLocation), true)

	fileLogger := slog.New(slog.NewJSONHandler(customlog.RotatingLog(configDS.LogLocation), nil))

	customlog.InfoFile(fileLogger, appconstants.STARTUP_MSG)
	customlog.InfoFile(fileLogger, "Config Values for this run ...")
	customlog.InfoFile(fileLogger, fmt.Sprintf("%+v", configDS))

	//Setup PANIC trap
	defer func() {
		r := recover()

		if r != nil {
			customlog.ErrorConsole(consoleLogger, r.(error).Error())
		}
	}()

	//Load simulator configuration data
	err = fileLoader(&configDS, &locationsDS, &routesDS, &classOfServiceDS, &qubzNameDS, &sensorTypeDS, &shipmentTypeDS, &transportModeDS, &sensorRangeDS, &sensorExceptionDS, consoleLogger, fileLogger)

	if err != nil {
		//Exit with Failure message
		customlog.ErrorAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_FAILED_MSG)
		return
	} else {
		fmt.Println(sensorExceptionDS)
	}

	//Determine the number of Qubz to simulate
	var qubz_simulation_count int = 0

	if configDS.QubzCount > len(qubzNameDS) || configDS.QubzCount == 0 {
		//Not enough Qubz names - adjust the count of Qubz to the max available OR no limit specified so just load all available QUBZ names
		qubz_simulation_count = len(qubzNameDS)

	} else if configDS.QubzCount < 0 {
		//Invalid value - notify and bail out
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Invalid qubzCount config value - cannot configure simulation")
		customlog.ErrorAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_FAILED_MSG)
		return

	} else {
		//Set number of Qubz to simulate to the number in the Config file
		qubz_simulation_count = configDS.QubzCount
	}

	//Connect to the Kafka output channel
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Initializing Kafka Connections ...", true)

	err = configureKafkaConnections(&skafkaConnections, configDS.QueueTopic, configDS.QueueEndpoint, consoleLogger, fileLogger)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return
	}

	//Allocate and initialize the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Initializing Qubz Matrix ...", true)
	currentQubzMatrix := make([]datamodels.QubzMatrix, qubz_simulation_count)

	err = initializeQubzMatrix(qubz_simulation_count, &qubzNameDS, &currentQubzMatrix, consoleLogger, fileLogger, &configDS)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Qubz Matrix Initialized", true)

	//Assign routes to Qubz in the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Assigning Routes to Qubz in Qubz Matrix ...", true)

	err = assignQubzRoutes(qubz_simulation_count, &routesDS, &currentQubzMatrix, consoleLogger, fileLogger, &configDS)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Route Assignment Complete", true)

	//Assign shipment types to Qubz in the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Assigning Shipment Types to Qubz in Qubz Matrix ...", true)

	err = assignQubzShipmentTypes(qubz_simulation_count, &shipmentTypeDS, &currentQubzMatrix, consoleLogger, fileLogger, &configDS)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Shipment Type Assignment Complete", true)

	//Configure initial exception levels in the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Configuring Initial Exceptions in Qubz Matrix ...", true)

	initializeQubzMatrixExceptions(&currentQubzMatrix, consoleLogger, fileLogger)

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Exception Configuration Complete", true)

	//Configure initial sensor values in the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Configuring Initial Sensor Values in Qubz Matrix ...", true)

	initializeQubzMatrixSensors(&currentQubzMatrix, &sensorRangeDS, consoleLogger, fileLogger)

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Qubz Matrix Sensor Configuration Complete", true)

	//Start simulation run
	//Allocate and initialize the Qubz Matrix
	customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Begin Simulation (%d cycles)", configDS.EventCycleCount), true)

	for i := 0; i < configDS.EventCycleCount; i++ {

		cycleNumber := i + 1

		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Capture Current Sensor State to Prior State Buffer", cycleNumber, configDS.EventCycleCount))
		priorQubzMatrix := make([]datamodels.QubzMatrix, len(currentQubzMatrix))
		copy(priorQubzMatrix, currentQubzMatrix)
		datautil.AssignEventStatesAll(&priorQubzMatrix, appconstants.SENSOR_STATE_PREVIOUS)

		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Assign New Exception States", cycleNumber, configDS.EventCycleCount))
		assignExceptions(&currentQubzMatrix, &sensorExceptionDS, consoleLogger, fileLogger)

		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Calculate New Sensor State", cycleNumber, configDS.EventCycleCount))
		updateQubzMatrixSensors(&currentQubzMatrix, &sensorRangeDS, &sensorExceptionDS, consoleLogger, fileLogger)
		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Sensor State Updated ... Ready for Transmission", cycleNumber, configDS.EventCycleCount))

		//fmt.Println(currentQubzMatrix)

		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Transmitting events to output channel", cycleNumber, configDS.EventCycleCount))
		err = eventemitter.TransmitEvents(&currentQubzMatrix, &priorQubzMatrix, &configDS, consoleLogger, fileLogger)

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, "Simulation Run Terminated!")
			return
		}

		customlog.CalloutAllChannels(consoleLogger, fileLogger, fmt.Sprintf("(Cycle %d of %d) Complete!", cycleNumber, configDS.EventCycleCount))

		//Check to see if there is a pause before the next cycle
		if configDS.EventInterval > 0 && i != (configDS.EventCycleCount-1) {
			customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Pausing for %d seconds before next cycle ...", configDS.EventInterval), false)
			time.Sleep(time.Duration(configDS.EventInterval) * time.Second)
		}

	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("%d cycles Complete ... Shutting Down", configDS.EventCycleCount), true)

	//Close the Kafka connections
	customlog.InfoAllChannels(consoleLogger, fileLogger, "Closing Kafka Connections ....", true)
	err = closeKafkaConnections(&skafkaConnections, consoleLogger, fileLogger)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Simulation Run Terminated!")
		return
	} else {
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Kafka Connections Closed", true)
	}

	//Exit with a success message
	customlog.InfoAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_COMPLETE_MSG, true)

}

func closeKafkaConnections(kafkaConnections *[]clientmodels.S_kafkaConnection, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {

	//Cycle through the list of Kafka connections and close each of them

	for _, x := range *kafkaConnections {

		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Closing connection for sensor \"%s\"", x.Key), false)

		err := x.Connection.Close()

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Error closing connection for sensor \"%s\" : %s", x.Key, err.Error()))
		}
	}

	return nil
}

func configureKafkaConnections(kafkaConnections *[]clientmodels.S_kafkaConnection, topicName string, brokerAddress string, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//Configure the connections in the passed connection Struct reference. TODO: Make this less "hard-coded" in next release

	var errsOccurred bool = false

	for i := 0; i < 13; i++ {

		//Set connection metadata
		switch i {
		case 0:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_ALTIMETER)
			(*kafkaConnections)[i].Partition = 0
			(*kafkaConnections)[i].Topic = topicName

		case 1:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_BATTERY)
			(*kafkaConnections)[i].Partition = 1
			(*kafkaConnections)[i].Topic = topicName

		case 2:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_COMPUTE)
			(*kafkaConnections)[i].Partition = 2
			(*kafkaConnections)[i].Topic = topicName
		case 3:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_GEIGER)
			(*kafkaConnections)[i].Partition = 3
			(*kafkaConnections)[i].Topic = topicName
		case 4:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_GPS)
			(*kafkaConnections)[i].Partition = 4
			(*kafkaConnections)[i].Topic = topicName
		case 5:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_GYRO)
			(*kafkaConnections)[i].Partition = 5
			(*kafkaConnections)[i].Topic = topicName
		case 6:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_LOCK)
			(*kafkaConnections)[i].Partition = 6
			(*kafkaConnections)[i].Topic = topicName
		case 7:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_MOTION)
			(*kafkaConnections)[i].Partition = 7
			(*kafkaConnections)[i].Topic = topicName
		case 8:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_RADIO)
			(*kafkaConnections)[i].Partition = 8
			(*kafkaConnections)[i].Topic = topicName
		case 9:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_SEAL)
			(*kafkaConnections)[i].Partition = 9
			(*kafkaConnections)[i].Topic = topicName
		case 10:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_SPECTROMETER)
			(*kafkaConnections)[i].Partition = 10
			(*kafkaConnections)[i].Topic = topicName
		case 11:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_TEMPHUMIDITY)
			(*kafkaConnections)[i].Partition = 11
			(*kafkaConnections)[i].Topic = topicName
		case 12:
			(*kafkaConnections)[i].Key = strings.ToLower(appconstants.SENSOR_TYPE_FIRE)
			(*kafkaConnections)[i].Partition = 12
			(*kafkaConnections)[i].Topic = topicName

		}

		//Get the connection
		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Configuring Kafka Connection for sensor \"%s\"", (*kafkaConnections)[i].Key), false)
		err := getKafkaConnection(kafkaConnections, i, brokerAddress)

		if err != nil {
			customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
			errsOccurred = true

		} else {
			customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Connection configured - Connected to topic \"%s\" on partition %d", (*kafkaConnections)[i].Topic, (*kafkaConnections)[i].Partition), false)
		}
	}

	if errsOccurred {
		return fmt.Errorf("fatal error(s) occurred configuring Kafka connections")
	} else {
		return nil
	}

}

func getKafkaConnection(kafkaConnections *[]clientmodels.S_kafkaConnection, connectionIndex int, brokerAddress string) error {

	conn, err := svcclient.GetSkafkaConnection((*kafkaConnections)[connectionIndex].Topic, brokerAddress, (*kafkaConnections)[connectionIndex].Partition)

	if err != nil {
		//Error - return it back to the caller
		return err
	}

	(*kafkaConnections)[connectionIndex].Connection = conn

	return nil
}
