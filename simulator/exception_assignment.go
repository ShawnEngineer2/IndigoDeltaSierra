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
		customlog.CalloutConsole(consoleLogger, "Get New Exception Severity Index")
		randomIndex := randomgen.RandomInt(1, len(distroMatrix))
		customlog.CalloutConsole(consoleLogger, fmt.Sprintf("Distro Matrix Len : %d | Severity Index: %d", len(distroMatrix), randomIndex))
		newExceptionSeverity := distroMatrix[randomIndex]

		if newExceptionSeverity != appconstants.SENSOR_EXCEPTION_SEVERITY_NONE {

			customlog.CalloutConsole(consoleLogger, "Get New Exception")
			newException := getException(&sevCritical, &sevHigh, &sevLow, newExceptionSeverity)

			//Compare the returned exception severity to what is currently assigned
			//for the current Qubz unit and assign a new Exception as needed
			customlog.CalloutConsole(consoleLogger, "Check Severity Level on Qubz Unit")

			if x.ExceptionSeverity == appconstants.SENSOR_EXCEPTION_SEVERITY_NONE {
				//Assign a new exception to this Qubz unit
				customlog.GreenlighAllChannels(consoleLogger, fileLogger, fmt.Sprintf("New Exception Assigned to Qubz Unit %s : Exception %s : Severity %d", x.QubzName, newException.ExceptionDesc, newException.SeverityLevel))
				assignException(qubzMatrix, newException, i)

			} else if x.ExceptionSeverity < newExceptionSeverity {
				//Upgrade this Qubz unit to a higher Severity
				customlog.GreenlighAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Exception level upgraded for Qubz Unit %s : Exception %s : Severity %d", x.QubzName, newException.ExceptionDesc, newException.SeverityLevel))
				assignException(qubzMatrix, newException, i)

			} else {
				//Notify that an exception was generated but not assigned
				customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Exception generated but not assigned to Qubz Unit %s : Exception Severity was %d : Qubz Unit Existing Exception Severity %d", x.QubzName, x.ExceptionSeverity, newException.SeverityLevel), true)

			}
		}
	}

}

func assignException(qubzMatrix *[]datamodels.QubzMatrix, newException datamodels.QubzException, matrixIndex int) {
	//Assign the passed exception to the passed index in the Qubz Matrix
	(*qubzMatrix)[matrixIndex].ExceptionAssignment = newException.ExceptionId
	(*qubzMatrix)[matrixIndex].ExceptionSeverity = newException.SeverityLevel
	(*qubzMatrix)[matrixIndex].ExceptionType = newException.ExceptionType
	(*qubzMatrix)[matrixIndex].ExceptionIntervalBoundary = newException.IntermittencyInterval
	(*qubzMatrix)[matrixIndex].CurrentExceptionInterval = 1

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
