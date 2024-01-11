package SvcClient

import (
	"fmt"
	"strconv"
)

func GetRandomNumbers(numValues int, numMin int, numMax int) {

	const baseSvcUrl string = "https://www.random.org/integersss"
	const numValueBoundary int = 1000

	//Validate Number of Values input
	if numValues < 1 {
		fmt.Println("Must request at least one random number")
	} else if numValues > numValueBoundary {
		msg := "Cannot request more than " + strconv.Itoa(numValueBoundary) + " numbers per service call"
		fmt.Println(msg)
	}

	//Validate Min and Max values
	if numMin < 0 {
		fmt.Println("Minimum value must be at least 0")
	} else if numMax < 1 {
		fmt.Println("Maximum value must be at least 1")
	}

	//Build Service URL
	svcUrl := baseSvcUrl + "/?num=" + strconv.Itoa(numValues) + "&min=" + strconv.Itoa(numMin) + "&max=" + strconv.Itoa(numMax) + "&col=" + strconv.Itoa(numValues) + "&base=10&format=plain&rnd=new"

	//Populate required headers
	var svcHeaders = []header{
		{
			headerName:  "User-Agent",
			headerValue: "shawn.engineer2@gmail.com",
		},
	}

	//Use "Get" from SvcBaseFunctions to retrieve results from the service
	httpCode, responseBytes := Get(svcUrl, "", svcHeaders)
	responseString := string(responseBytes)

	//Handle non-OK failure codes - taking the simple route here
	if httpCode != 200 && httpCode != 201 {
		msg := "Service Error Occurred : Status Code : " + strconv.Itoa(httpCode) + " Error Message : " + responseString
		fmt.Println(msg)
	}

	fmt.Println(responseString)
}
