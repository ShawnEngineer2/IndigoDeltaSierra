package svcclient

import (
	"errors"
	"log/slog"
	"strconv"
	"strings"
)

func GetRandomNumbers(numValues int, numMin int, numMax int, logger *slog.Logger, userAgent string) []string {

	const baseSvcUrl string = "https://www.random.org/integers"
	const numValueBoundary int = 1000

	//Validate Number of Values input
	if numValues < 1 {
		logger.Error("Must request at least one random number")
		return nil

	} else if numValues > numValueBoundary {
		msg := "Cannot request more than " + strconv.Itoa(numValueBoundary) + " numbers per service call"
		logger.Error(msg)
		return nil
	}

	//Validate Min and Max values
	if numMin < 0 {
		logger.Error("Minimum value must be at least 0")
		return nil

	} else if numMax < 1 {
		logger.Error("Maximum value must be at least 1")
		return nil

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
	httpCode, responseBytes := Get(svcUrl, "", svcHeaders, logger)
	responseString := string(responseBytes)

	//Handle non-OK failure codes - taking the simple route here
	if httpCode != 200 && httpCode != 201 {
		msg := "Service Error Occurred : Status Code : " + strconv.Itoa(httpCode) + " Error Message : " + responseString
		logger.Error(msg)
		return nil
	}

	//Return the body data
	payload := strings.Split(responseString, "\t")
	return payload

}

func CheckQuotaExceeded(fileLogger *slog.Logger, userAgent string) (bool, error) {
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
	httpCode, responseBytes := Get(svcUrl, "", svcHeaders, fileLogger)
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
