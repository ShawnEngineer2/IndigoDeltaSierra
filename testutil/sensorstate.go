package testutil

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"indigodeltasierra/randomgen"
	"log/slog"
)

//This package tests to make sure sensor states are properly calculated and set

func TestSensorStateUpdate() {

	//Set up console logger
	consoleLogger := slog.Default()
	customlog.InfoConsole(consoleLogger, "Beginning Sensor Value test ...", true)

	//Load the Config
	configDS := datamodels.Config{}
	err := datautil.LoadDataFile(&configDS, "Configuration", appconstants.CONFIG_FILE_PATH, consoleLogger, nil)

	if err != nil {
		customlog.ErrorConsole(consoleLogger, err.Error())
		return
	}

	//Setup logfile logger
	customlog.InfoConsole(consoleLogger, fmt.Sprintf("Configuring File Logger to output path %s ...", configDS.LogLocation), true)

	fileLogger := slog.New(slog.NewJSONHandler(customlog.RotatingLog(configDS.LogLocation), nil))

	customlog.InfoFile(fileLogger, appconstants.STARTUP_MSG)
	customlog.InfoFile(fileLogger, "Config Values for this run ...")
	customlog.InfoFile(fileLogger, fmt.Sprintf("%+v", configDS))

	//Load the list of Sensor Range values
	sensorRangeDS := make([]datamodels.SensorRange, 1)
	err = datautil.LoadDataFile(&sensorRangeDS, "Sensor Range", configDS.SensorRangesFile, consoleLogger, fileLogger)

	if err != nil {
		customlog.ErrorConsole(consoleLogger, err.Error())
		return
	}

	//Initialize Qubz State structure
	qubzStateDS := datamodels.QubzState{}
	qubzStateDS.SensorDataPoints = make([]datamodels.QubzSensorDataPoint, len(sensorRangeDS))

	qubzStateDS.ExceptionAssignment = -1
	qubzStateDS.ExceptionSeverity = 0
	qubzStateDS.ExceptionType = 0
	qubzStateDS.ExceptionIntervalBoundary = 0
	qubzStateDS.CurrentExceptionInterval = 0
	qubzStateDS.QubzID = 1
	qubzStateDS.QubzName = "ParseyMcParseFace"
	qubzStateDS.RouteAssignment = 1
	qubzStateDS.ShipmentType = 3
	qubzStateDS.TransportMode = 5

	//Working struct

	for i, x := range sensorRangeDS {

		//Fill in info
		qubzStateDS.SensorDataPoints[i].DataPointTypeId = x.DataPointTypeId
		qubzStateDS.SensorDataPoints[i].SensorDataPointId = x.SensorDataPointId

		//Calculate the value based on the type of data point
		switch x.DataPointTypeId {
		case appconstants.DATA_POINT_TYPE_RANGED:
			fmt.Println(x.SensorDataPoint, x.SensorDataPointId)
			qubzStateDS.SensorDataPoints[i].DataPointValue = randomgen.RandomIFloat(x.NominalMin, x.NominalMax, x.NumberScale)

		case appconstants.DATA_POINT_TYPE_BOOLEAN:
			qubzStateDS.SensorDataPoints[i].DataPointValue = x.NominalMin

		default:
			qubzStateDS.SensorDataPoints[i].DataPointValue = 0.00
		}

	}

	fmt.Println(qubzStateDS)

	customlog.InfoConsole(consoleLogger, "Sensor Value Test Complete", true)
}
