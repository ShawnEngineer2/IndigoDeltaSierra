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

	//If Random Service is active then assign ship types using selections from the service
	if config.InternalRandomizer == 0 {

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
	}

	//If Random Service is bypassed, assign ship type via a round robin algorithm
	if config.InternalRandomizer == 1 {

		customlog.InfoAllChannels(consoleLogger, fileLogger, "Assigning Shipment Types via Round Robin Algorithm ...", false)

		shipTypeIndex := 0

		for i := 0; i < numQubz; i++ {

			(*qubzMatrix)[i].ShipmentType = (*qubzShipmentTypes)[shipTypeIndex].ShipmentTypeId

			//Advance and adjust the route pointer as needed
			shipTypeIndex++

			if shipTypeIndex == len(*qubzShipmentTypes) {
				//Reset the route pointer to the start of the Route list
				shipTypeIndex = 0
			}

		}

	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Shipment Types mapped to Qubz Matrix", false)

	return nil

}
