package SvcClient

import (
	"fmt"
	"io"
	"net/http"
)

type header struct {
	headerName  string
	headerValue string
}

func Get(svcEndpoint string, body string, headers []header) (int, []byte) {
	fmt.Println("Made it to GET")

	fmt.Println("Configuring Call ...")

	client := &http.Client{}

	//Add Body
	if len(body) > 0 {
		fmt.Println("Adding Body")
	}

	req, err := http.NewRequest("GET", svcEndpoint, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	//Add headers here
	if headers != nil {
		fmt.Println("Adding Headers")

		for _, currHeader := range headers {
			msg := "Adding Header " + currHeader.headerName + " with Value " + currHeader.headerValue
			fmt.Println(msg)
			req.Header.Add(currHeader.headerName, currHeader.headerValue)
		}
	}

	fmt.Println("Making Call ...")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	return resp.StatusCode, bodyBytes

}
