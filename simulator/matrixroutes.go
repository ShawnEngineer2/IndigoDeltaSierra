package simulator

import (
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/svcclient"
	"log/slog"
	"strconv"
	"strings"
)

func assignQubzRoutes(numQubz int, qubzRoutes *[]datamodels.Route, qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger, config *datamodels.Config) error {
	//This function assigns routes to the Qubz in the Qubz Matrix

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Selecting Route Assignments ...", false)

	selectedRoutes, err := svcclient.GetRandomNumbers(numQubz, 0, (len(*qubzRoutes) - 1), consoleLogger, fileLogger, config.EmailAddress)

	if err != nil {
		return err
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Mapping Routes to Qubz Matrix ...", false)

	var selectedRouteIndex int = 0

	for i := 0; i < len(selectedRoutes); i++ {
		selectedRouteIndex, err = strconv.Atoi(strings.TrimSuffix(selectedRoutes[i], "\n"))

		if err != nil {
			return err
		}

		(*qubzMatrix)[i].RouteAssignment = (*qubzRoutes)[selectedRouteIndex].RouteID
		(*qubzMatrix)[i].TransportMode = (*qubzRoutes)[selectedRouteIndex].TransportModeID
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Routes mapped to Qubz Matrix", false)

	return nil

}
