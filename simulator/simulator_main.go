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

	fmt.Println(sensorRangeDS)

	//Exit with a success message
	customlog.InfoAllChannels(consoleLogger, fileLogger, appconstants.SIMULATION_COMPLETE_MSG, true)

}
