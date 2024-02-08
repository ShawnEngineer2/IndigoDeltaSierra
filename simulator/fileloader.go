package simulator

import (
	"errors"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func fileLoader(locationsDS *[]datamodels.Location, routesDS *[]datamodels.Route, classOfServiceDS *[]datamodels.ClassOfService, qubzNameDS *[]datamodels.Qubz, sensorTypeDS *[]datamodels.SensorType, shipmentTypeDS *[]datamodels.ShipmentType, transportModeDS *[]datamodels.TransportMode, sensorRangeDS *[]datamodels.SensorRange, consoleLogger *slog.Logger, fileLogger *slog.Logger) error {
	//This routine handles loading simulator data structs with information from data files

	const config_FAILED_MSG string = "Simulator Data Configuration Failed"

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Beginning Simulator Data Configuration ...", true)

	//Load Locations data
	err := datautil.LoadDataFile(&locationsDS, "Locations", appconstants.LOCATIONS_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load Route data
	err = datautil.LoadDataFile(&routesDS, "Routes", appconstants.ROUTES_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load Class of Service data
	err = datautil.LoadDataFile(&classOfServiceDS, "Class of Service", appconstants.CLASS_OF_SERVICE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Qubz names
	err = datautil.LoadDataFile(&qubzNameDS, "Qubz Names", appconstants.QUBZ_NAME_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Sensor Types
	err = datautil.LoadDataFile(&sensorTypeDS, "Sensor Types", appconstants.SENSOR_TYPE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Shipment Types
	err = datautil.LoadDataFile(&shipmentTypeDS, "Shipment Types", appconstants.SHIPMENT_TYPES_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Transport Modes
	err = datautil.LoadDataFile(&transportModeDS, "Transport Modes", appconstants.TRANSPORT_MODE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	//Load the list of Sensor Range values
	err = datautil.LoadDataFile(&sensorRangeDS, "Sensor Range", appconstants.SENSOR_RANGE_FILE_PATH, consoleLogger, fileLogger)

	if err != nil {
		return errors.New(config_FAILED_MSG)
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Simulator Data Configuration Complete", true)

	return nil
}
