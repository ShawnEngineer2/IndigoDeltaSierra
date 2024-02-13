package simulator

import (
	"fmt"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/svcclient"
	"log/slog"
	"strconv"
	"strings"
)

func initializeQubzMatrix(numQubz int, qubzNamesDS *[]datamodels.Qubz, qubzMatrix *[]datamodels.QubzMatrix, consoleLogger *slog.Logger, fileLogger *slog.Logger, config *datamodels.Config) error {
	//This function initializes the passed Qubz Matrix structure with Qubz that are in-scope for this simulation

	//If the number of requested Qubz equals the total number of Qubz in the qubzNames list,
	//just copy all ids into the Qubz matrix
	fmt.Println(len(*qubzNamesDS))

	if numQubz == len(*qubzNamesDS) {

		customlog.InfoAllChannels(consoleLogger, fileLogger, "Mapping All QubzIds into Qubz Matrix ...", false)

		for i, x := range *qubzNamesDS {
			(*qubzMatrix)[i].QubzID = x.QubzID
			(*qubzMatrix)[i].QubzName = x.QubzName
		}
	}

	//If the number of requested qubz is less than the total number of Qubz in qubzNames and Random Service isn't bypassed,
	//use the Random Number online service to select the Qubz to use
	if (numQubz < len(*qubzNamesDS)) && config.InternalRandomizer == 0 {
		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Using Random Service to selct %d QubzIds to map into Qubz Matrix ...", numQubz), false)
		selectedQubzIds, err := svcclient.GetRandomNumbers(numQubz, 0, (len(*qubzNamesDS) - 1), consoleLogger, fileLogger, config.EmailAddress)

		if err != nil {
			return err
		}

		customlog.InfoAllChannels(consoleLogger, fileLogger, "Mapping Selected QubzIds into Qubz Matrix ...", false)

		var selectedQubzIndex int = 0

		for i := 0; i < len(selectedQubzIds); i++ {
			selectedQubzIndex, err = strconv.Atoi(strings.TrimSuffix(selectedQubzIds[i], "\n"))

			if err != nil {
				return err
			}

			(*qubzMatrix)[i].QubzID = (*qubzNamesDS)[selectedQubzIndex].QubzID
			(*qubzMatrix)[i].QubzName = (*qubzNamesDS)[selectedQubzIndex].QubzName
		}

	}

	//If the number of requested qubz is less than the total number of Qubz in qubzNames and we're bypassing Random Service,
	//use a simple algorithm to select the Qubz to use
	if (numQubz < len(*qubzNamesDS)) && config.InternalRandomizer == 1 {

		customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Mapping first %d QubzIds into Qubz Matrix ...", numQubz), false)

		for i := 0; i < numQubz; i++ {

			(*qubzMatrix)[i].QubzID = (*qubzNamesDS)[i].QubzID
			(*qubzMatrix)[i].QubzName = (*qubzNamesDS)[i].QubzName
		}

	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Qubz Matrix initialization complete - %d QubzIds mapped into Qubz Matrix", len(*qubzMatrix)), false)

	return nil

}