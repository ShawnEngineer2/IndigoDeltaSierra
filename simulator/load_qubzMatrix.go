package simulator

/*

var max_qubz_count int = 0

	//Load Qubz Names and IDs from the Qubz Name File
	qubznames := []datamodels.Qubz{}

	if !sysfile.LoadFileToStruct(config.QubzNameFile, &qubznames) {
		consoleLogger.Error(fmt.Sprintf("Could not load Qubz Names from file %s ... startup terminated", config.QubzNameFile))
		//return error
	} else {
		max_qubz_count = len(qubznames)
		consoleLogger.Info(fmt.Sprintf("Qubz Names loaded from file %s ... %d names loaded", config.QubzNameFile, max_qubz_count))
	}

	//Use the loaded names list to source a random set of Qubz names
	qubzmatrix := make([]datamodels.QubzMatrix, config.QubzCount)

	if config.QubzCount == 0 || config.QubzCount == max_qubz_count {
		//Just load'em all into the Qubz Matrix
		fileLogger.Info("Loading All Qubz Names ...")

		qubzmatrix = make([]datamodels.QubzMatrix, max_qubz_count)

		for i, x := range qubznames {
			qubzmatrix[i].QubzID = x.QubzID
			qubzmatrix[i].QubzName = x.QubzName
		}

		//fmt.Println(qubzmatrix[10])

	} else if config.QubzCount < max_qubz_count {
		//Grab a random set of numbers of QubzCount between 1 and max qubz count and load those names
		fileLogger.Info("Loading Random Qubz Names ...")

		fmt.Println(svcclient.GetRandomNumbers(config.QubzCount, 0, (max_qubz_count - 1), fileLogger, config.EmailAddress))

	} else if config.QubzCount > max_qubz_count {
		//Too many Qubz requested - throw an error
		consoleLogger.Error(fmt.Sprintf("Not enough names in %s file (%d) to satisfy requested number of Qubz (%d) ... startup terminated", config.QubzNameFile, len(qubznames), config.QubzCount))
		//return error
	}

	err := datautil.InitializeQubzMatrix(&qubzmatrix)


*/
