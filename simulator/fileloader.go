package simulator

import (
	"errors"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func fileLoader(config *datamodels.Config, locationsDS *[]datamodels.Location, routesDS *[]datamodels.Route, classOfServiceDS *[]datamodels.ClassOfService, qubzNameDS *[]datamodels.Qubz, sensorTypeDS *[]datamodels.SensorType, shipmentTypeDS *[]datamodels.ShipmentType, transportModeDS *[]datamodels.TransportMode, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//This routine handles loading simulator data structs with information from data files

	const config_FAILED_MSG string = "simulator data configuration failed"

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Beginning Simulator Data Configuration ...", true)

	//Load Locations data
	err := datautil.LoadDataFile(locationsDS, "Locations", config.LocationsFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load Route data
	err = datautil.LoadDataFile(routesDS, "Routes", config.RoutesFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load Class of Service data
	err = datautil.LoadDataFile(classOfServiceDS, "Class of Service", config.ClassOfServiceFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Qubz names
	err = datautil.LoadDataFile(qubzNameDS, "Qubz Names", config.QubzNameFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Sensor Types
	err = datautil.LoadDataFile(sensorTypeDS, "Sensor Types", config.SensorTypesFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Shipment Types
	err = datautil.LoadDataFile(shipmentTypeDS, "Shipment Types", config.ShipmentTypesFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Transport Modes
	err = datautil.LoadDataFile(transportModeDS, "Transport Modes", config.TransportModesFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Sensor Range values
	err = datautil.LoadDataFile(sensorRangeDS, "Sensor Range", config.SensorRangesFile, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Simulator Data Configuration Complete", true)

	return nil
}
