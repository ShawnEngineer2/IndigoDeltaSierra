package svcclient

import (
	"fmt"
	"indigodeltasierra/customlog"
	"io"
	"log/slog"
	"net/http"
)

type header struct {
	headerName  string
	headerValue string
}

func Get(svcEndpoint string, body string, headers []header, consoleLogger *slog.Logger, fileLogger *slog.Logger) (int, []byte) {

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Configuring Service Call", false)

	client := &http.Client{}

	//Add Body
	if len(body) > 0 {
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Adding Service Request Body ...", false)
	}

	req, err := http.NewRequest("GET", svcEndpoint, nil)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
	}

	//Add headers here
	if headers != nil {
		customlog.InfoAllChannels(consoleLogger, fileLogger, "Adding Service Request Headers ...", false)

		for _, currHeader := range headers {
			customlog.InfoAllChannels(consoleLogger, fileLogger, fmt.Sprintf("Adding Header %s with Value %s ...", currHeader.headerName, currHeader.headerValue), false)
			req.Header.Add(currHeader.headerName, currHeader.headerValue)
		}
	}

	customlog.InfoAllChannels(consoleLogger, fileLogger, "Making Service Call ...", false)

	resp, err := client.Do(req)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		customlog.ErrorAllChannels(consoleLogger, fileLogger, err.Error())
	}

	return resp.StatusCode, bodyBytes

}
