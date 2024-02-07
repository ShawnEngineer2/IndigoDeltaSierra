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

	//Load Locations data
	err = datautil.LoadDataFile(&locationsDS, "Locations", appconstants.LOCATIONS_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load Route data
	err = datautil.LoadDataFile(&routesDS, "Routes", appconstants.ROUTES_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load Class of Service data
	err = datautil.LoadDataFile(&classOfServiceDS, "Class of Service", appconstants.CLASS_OF_SERVICE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load the list of Qubz names
	err = datautil.LoadDataFile(&qubzNameDS, "Qubz Names", appconstants.QUBZ_NAME_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load the list of Sensor Types
	err = datautil.LoadDataFile(&sensorTypeDS, "Sensor Types", appconstants.SENSOR_TYPE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load the list of Shipment Types
	err = datautil.LoadDataFile(&shipmentTypeDS, "Shipment Types", appconstants.SHIPMENT_TYPES_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	//Load the list of Transport Modes
	err = datautil.LoadDataFile(&transportModeDS, "Transport Modes", appconstants.TRANSPORT_MODE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return
	}

	fmt.Println(transportModeDS)

	//Exit with a CLEAN (no errors) code
	customlog.InfoAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_COMPLETE_MSG, true)

}
