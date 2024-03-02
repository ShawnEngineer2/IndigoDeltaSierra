package datautil

import (
	"errors"
	"fmt"
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/randomgen"
)

func GetExceptionList(exceptionDS *[]datamodels.QubzException, severityLevel int) ([]datamodels.QubzException, error) {
	//This routine returns a slice of Exception structs matching the passed Severity level
	exceptionList := []datamodels.QubzException{}
	foundCount := 0

	//Walk through the list of exception definitions and return a list of those matching the passed severity
	for _, x := range *exceptionDS {

		if x.SeverityLevel == severityLevel {
			//Found one - add it to the slice
			exceptionList = append(exceptionList, x)
			foundCount++
		}
	}

	//Check to see if we found anything - if we didn't, return an error
	if foundCount == 0 {
		//return an error
		return nil, errors.New(fmt.Sprintf("no exceptions found with Severity Level %d", severityLevel))
	}

	//Return the results
	return exceptionList, nil
}

func GetRandomException(exceptionDS *[]datamodels.QubzException) datamodels.QubzException {
	//This routine returns a single random exception from the passed list

	newException := datamodels.QubzException{}
	exceptionIndexBoundary := len(*exceptionDS)

	if exceptionIndexBoundary == 1 {
		//Only one exception - just return it as the new exception
		newException = (*exceptionDS)[0]
	} else if exceptionIndexBoundary > 0 {
		randomindex := randomgen.RandomInt(1, exceptionIndexBoundary)
		newException = (*exceptionDS)[randomindex-1]
	} else {
		//Empty list passed - throw error
		customerror.CheckAndPanic(errors.New("cannot retrieve random exception from an empty exception list"))
	}

	return newException
}

func GetSingleException(exceptionDS *[]datamodels.QubzException, exceptionId int) (datamodels.QubzException, error) {
	//This routine returns the exception in the passed list matching the passed exception Id

	for _, x := range *exceptionDS {

		//Compare the passed exception id to the exception id of the qubz exception
		//and if match pass back to the caller
		if x.ExceptionId == exceptionId {
			return x, nil
		}
	}

	//If you get this far, you didn't find anything - return an error
	err := fmt.Errorf("cannot locate an exception defintion for exception id %d", exceptionId)
	emptyException := datamodels.QubzException{}

	return emptyException, err
}
