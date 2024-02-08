package simulator

import (
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/svcclient"
	"log/slog"
	"strconv"
	"strings"
)

func assignQubzShipmentTypes(numQubz int, qubzShipmentTypes *[]datamodels.ShipmentType, qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger, config *datamodels.Config) error {
	//This function assigns shipment types to the Qubz in the Qubz Matrix

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Selecting Shipment Type Assignments ...", false)

	selectedShipTypes, err := svcclient.GetRandomNumbers(numQubz, 0, (len(*qubzShipmentTypes) - 1), consoleLogger, fileLogger, config.EmailAddress)

	if err != nil {
		return err
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Mapping Shipment Types to Qubz Matrix ...", false)

	var selectedShipTypeIndex int = 0

	for i := 0; i < len(selectedShipTypes); i++ {
		selectedShipTypeIndex, err = strconv.Atoi(strings.TrimSuffix(selectedShipTypes[i], "\n"))

		if err != nil {
			return err
		}

		(*qubzMatrix)[i].ShipmentType = (*qubzShipmentTypes)[selectedShipTypeIndex].ShipmentTypeId
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Shipment Types mapped to Qubz Matrix", false)

	return nil

}
