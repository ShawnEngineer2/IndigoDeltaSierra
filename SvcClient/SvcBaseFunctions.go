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

func Get(svcEndpoint string, body string, headers []header) []byte {
	fmt.Println("Made it to GET")

	fmt.Println("Configuring Call ...")

	client := &http.Client{}

	//Add headers here

	//Add Body

	req, err := http.NewRequest("GET", svcEndpoint, nil)

	if err != nil {
		fmt.Println(err.Error())
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

	return bodyBytes

}
