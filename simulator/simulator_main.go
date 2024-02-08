package simulator

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
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

	//Set up console logger
	consoleLogger := slog.Default()
	customlog.InfoConsole(consoleLogger, appconstants.STARTUP_MSG, true)

	//Load the Config file
	err := datautil.LoadDataFile(&configDS, "Configuration", appconstants.CONFIG_FILE_PATH, consoleLogger, nil)

	if err != nil {
		return
	}

	//Setup logfile logger
	customlog.InfoConsole(consoleLogger, "Configuring File Logger to output path "+configDS.LogLocation+" ...", true)

	fileLogger := slog.New(slog.NewJSONHandler(customlog.RotatingLog(configDS.LogLocation), nil))

	customlog.InfoFile(fileLogger, appconstants.STARTUP_MSG)
	customlog.InfoFile(fileLogger, "Config Values for this run ...")
	customlog.InfoFile(fileLogger, fmt.Sprintf("%+v", configDS))

	//Load simulator configuration data
	err = fileLoader(&locationsDS, &routesDS, &classOfServiceDS, &qubzNameDS, &sensorTypeDS, &shipmentTypeDS, &transportModeDS, &sensorRangeDS, consoleLogger, fileLogger)

	if err != nil {
		//Exit with Failure message
		customlog.ErrorAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_FAILED_MSG)
		return
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

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Qubz Matrix Sensor Configuration Complete", true)

	//Exit with a success message
	customlog.InfoAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_COMPLETE_MSG, true)

}
