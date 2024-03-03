package simulator

import (
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func logExceptionsAndCollectMetrics(currentQubzMatrix *[]datamodels.QubzMatrix, exceptionDS *[]datamodels.QubzException, numCycles int, numSensors int, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.ProcessingMetrics {
	//This routine pushes a list of the assigned exceptions to the log, then returns a populated Processing Metrics struct

	processingMetrics := datamodels.ProcessingMetrics{}

	processingMetrics.ActualEventCount = 0
	processingMetrics.EstimatedEventCount = 0
	processingMetrics.Sev01ExceptionCount = 0
	processingMetrics.Sev02ExceptionCount = 0
	processingMetrics.Sev03ExceptionCount = 0
	processingMetrics.TotalQubzCount = 0

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Calculating Metrics ....", false)

	customlog.InfoFile(fileLogger, "=================================================================================================")
	customlog.InfoFile(fileLogger, "==========                      ASSIGNED EXCEPTION LIST                        ==================")
	customlog.InfoFile(fileLogger, "=================================================================================================")

	for _, x := range *currentQubzMatrix {

		//Increment the Total Qubz Count
		processingMetrics.TotalQubzCount++

		//Check for exceptions and process as necessary
		if x.ExceptionType != appconstants.SENSOR_EXCEPTION_TYPE_NONE {

			//Update the Exception count
			switch x.ExceptionSeverity {
			case 1:
				processingMetrics.Sev01ExceptionCount++
			case 2:
				processingMetrics.Sev02ExceptionCount++
			case 3:
				processingMetrics.Sev03ExceptionCount++
			}

			//Retrieve the exception and print it to the Log
			var exceptionDesc string = ""
			exceptionDef, err := datautil.GetSingleException(exceptionDS, x.ExceptionAssignment)

			if err != nil {
				customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
				exceptionDesc = "UNDEFINED"
			} else {
				exceptionDesc = exceptionDef.ExceptionDesc
			}

			customlog.InfoFile(fileLogger, fmt.Sprintf("Exception %s : Severity %d assigned to Qubz Unit \"%s\"", exceptionDesc, x.ExceptionSeverity, x.QubzName))

		}
	}

	//Calculate final metrics
	processingMetrics.TotalExceptionCount = processingMetrics.Sev01ExceptionCount + processingMetrics.Sev02ExceptionCount + processingMetrics.Sev03ExceptionCount
	processingMetrics.EstimatedEventCount = processingMetrics.TotalQubzCount * numSensors * numCycles
	processingMetrics.ActualEventCount = processingMetrics.EstimatedEventCount

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Metrics Complete", false)

	customlog.InfoFile(fileLogger, "=================================================================================================")
	customlog.InfoFile(fileLogger, "===========                      END OF EXCEPTION LIST                        ===================")
	customlog.InfoFile(fileLogger, "=================================================================================================")

	return processingMetrics
}
