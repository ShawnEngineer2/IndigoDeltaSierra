package svcclient

import (
	"errors"
	"fmt"
	"indigodeltasierra/customlog"
	"log/slog"
	"strconv"
	"strings"
)

func GetRandomNumbers(numValues int, numMin int, numMax int, consoleLogger *slog.Logger, fileLogger *slog.Logger, userAgent string) ([]string, error) {

	const baseSvcUrl string = "https://www.random.org/integers"
	const numValueBoundary int = 10000
	const SVC_FAIL_MSG string = "call to random number service failed"

	//Validate Number of Values input
	if numValues < 1 {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Must request at least one random number")
		return nil, errors.New(SVC_FAIL_MSG)

	} else if numValues > numValueBoundary {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Cannot request more than %d numbers per service call", numValueBoundary))
		return nil, errors.New(SVC_FAIL_MSG)
	}

	//Validate Min and Max values
	if numMin < 0 {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Minimum value must be at least 0")
		return nil, errors.New(SVC_FAIL_MSG)

	} else if numMax < 1 {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, "Maximum value must be at least 1")
		return nil, errors.New(SVC_FAIL_MSG)

	}

	//Build Service URL
	svcUrl := baseSvcUrl + "/?num=" + strconv.Itoa(numValues) + "&min=" + strconv.Itoa(numMin) + "&max=" + strconv.Itoa(numMax) + "&col=" + strconv.Itoa(numValues) + "&base=10&format=plain&rnd=new"

	//Populate required headers
	var svcHeaders = []header{
		{
			headerName:  "User-Agent",
			headerValue: userAgent,
		},
	}

	//Use "Get" from SvcBaseFunctions to retrieve results from the service
	httpCode, responseBytes := Get(svcUrl, "", svcHeaders, consoleLogger, fileLogger)
	responseString := string(responseBytes)

	//Handle non-OK failure codes - taking the simple route here
	if httpCode != 200 && httpCode != 201 {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Service Error Occurred : Status Code : %d Error Message : %s", httpCode, responseString))
		return nil, errors.New(SVC_FAIL_MSG)
	}

	//Return the body data
	payload := strings.Split(responseString, "\t")
	return payload, nil

}

func CheckQuotaExceeded(consoleLogger *slog.Logger, fileLogger *slog.Logger, userAgent string) (bool, error) {
	//This function checks the random number service remaining bit quota and returns TRUE if available bits is less than the Bit Boundary
	const svcUrl string = "https://www.random.org/quota/?format=plain"
	const bitBoundary int = 1000

	//Populate required headers
	var svcHeaders = []header{
		{
			headerName:  "User-Agent",
			headerValue: userAgent,
		},
	}

	//Use "Get" from SvcBaseFunctions to retrieve results from the service
	httpCode, responseBytes := Get(svcUrl, "", svcHeaders, consoleLogger, fileLogger)
	responseString := string(responseBytes)

	//Handle non-OK failure codes - taking the simple route here
	if httpCode != 200 && httpCode != 201 {
		msg := "Service Error Occurred : Status Code : " + strconv.Itoa(httpCode) + " Error Message : " + responseString
		fileLogger.Error(msg)
		return false, errors.New("service error occurred")
	}

	//Process the returned response
	//----------------------------------
	quotaExceeded := false

	quotaInt, err := strconv.Atoi(strings.TrimRight(responseString, "\n"))

	if err != nil {
		//Log the error
		fileLogger.Error("Error Checking Random Service Quota: " + err.Error())
		return false, err
	}

	if quotaInt < bitBoundary {
		quotaExceeded = true
	}

	//Return the body data
	return quotaExceeded, nil

}
