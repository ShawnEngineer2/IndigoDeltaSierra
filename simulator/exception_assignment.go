package simulator

import (
	"errors"
	"fmt"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/customerror"
	"indigodeltasierra/customlog"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"indigodeltasierra/randomgen"
	"log/slog"
)

func assignExceptions(qubzMatrix *[]datamodels.QubzMatrix, exceptionDS *[]datamodels.QubzException, consoleLogger *slog.Logger, fileLogger *slog.Logger) {

	//Create individual exception lists by Exception Severity
	customlog.CalloutConsole(consoleLogger, "Divide Exceptions")
	sevCritical, sevHigh, sevLow, err := divideExceptionsBySeverity(exceptionDS, consoleLogger, fileLogger)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
	}

	//Prepare a distribution matrix that can be used for assigning exceptions to Qubz units
	customlog.CalloutConsole(consoleLogger, "Get Distro Matrix")
	distroMatrix := randomgen.CreateDistributionMatrix()

	//Walk through the Qubz Matrix and assign random exceptions
	for i, x := range *qubzMatrix {

		//Check to see if an exception is to be assigned for the current Qubz unit
		customlog.CalloutConsole(consoleLogger, "New Exception Severity")
		randomIndex := randomgen.RandomInt(1, len(distroMatrix))
		newExceptionSeverity := distroMatrix[randomIndex-1]

		if newExceptionSeverity != appconstants.SENSOR_EXCEPTION_SEVERITY_NONE {

			customlog.CalloutConsole(consoleLogger, "Check Severity Level on Qubz Unit")
			//Compare the returned exception severity to what is currently assigned
			//for the current Qubz unit
			if x.ExceptionSeverity == appconstants.SENSOR_EXCEPTION_SEVERITY_NONE || x.ExceptionSeverity < newExceptionSeverity {
				//Assign a new exception to this Qubz unit
				customlog.CalloutConsole(consoleLogger, "Get New Exception")
				newException := getException(&sevCritical, &sevHigh, &sevLow, newExceptionSeverity)

				(*qubzMatrix)[i].ExceptionAssignment = newException.ExceptionId
				(*qubzMatrix)[i].ExceptionSeverity = newException.SeverityLevel
				(*qubzMatrix)[i].ExceptionType = newException.ExceptionType
				(*qubzMatrix)[i].ExceptionIntervalBoundary = newException.IntermittencyInterval
				(*qubzMatrix)[i].CurrentExceptionInterval = 1

			}
		}
	}

}

func divideExceptionsBySeverity(exceptionDS *[]datamodels.QubzException, consoleLogger *slog.Logger, fileLogger *slog.Logger) ([]datamodels.QubzException, []datamodels.QubzException, []datamodels.QubzException, error) {

	//This routine creates 3 structs for Critical, High, and Low exceptions and returns them
	const ERROR_MSG string = "no exceptions found for severity %s"
	sevCritical, err := datautil.GetExceptionList(exceptionDS, appconstants.SENSOR_EXCEPTION_SEVERITY_CRITICAL)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return nil, nil, nil, errors.New(fmt.Sprintf(ERROR_MSG, "Critical"))
	}

	sevHigh, err := datautil.GetExceptionList(exceptionDS, appconstants.SENSOR_EXCEPTION_SEVERITY_HIGH)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return nil, nil, nil, errors.New(fmt.Sprintf(ERROR_MSG, "High"))
	}

	sevLow, err := datautil.GetExceptionList(exceptionDS, appconstants.SENSOR_EXCEPTION_SEVERITY_LOW)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
		return nil, nil, nil, errors.New(fmt.Sprintf(ERROR_MSG, "Low"))
	}

	return sevCritical, sevHigh, sevLow, nil

}

func getException(criticalExceptionDS *[]datamodels.QubzException, highExceptionDS *[]datamodels.QubzException, lowExceptionDS *[]datamodels.QubzException, severityLevel int) datamodels.QubzException {
	//Retrieve a random exception of the indicated severity

	newException := datamodels.QubzException{}

	switch severityLevel {
	case appconstants.SENSOR_EXCEPTION_SEVERITY_CRITICAL:
		newException = datautil.GetRandomException(criticalExceptionDS)

	case appconstants.SENSOR_EXCEPTION_SEVERITY_HIGH:
		newException = datautil.GetRandomException(highExceptionDS)

	case appconstants.SENSOR_EXCEPTION_SEVERITY_LOW:
		newException = datautil.GetRandomException(lowExceptionDS)

	default:
		//If you get this far, bail out
		msg := fmt.Sprintf("cannot assign exception - encountered invalid security level (%d)", severityLevel)
		customerror.CheckAndPanic(errors.New(msg))
	}

	return newException
}
