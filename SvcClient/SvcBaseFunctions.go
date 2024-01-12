package SvcClient

import (
	"io"
	"log/slog"
	"net/http"
)

type header struct {
	headerName  string
	headerValue string
}

func Get(svcEndpoint string, body string, headers []header, logger *slog.Logger) (int, []byte) {

	logger.Info("Configuring Call")

	client := &http.Client{}

	//Add Body
	if len(body) > 0 {
		logger.Info("Adding Body")
	}

	req, err := http.NewRequest("GET", svcEndpoint, nil)

	if err != nil {
		logger.Error(err.Error())
	}

	//Add headers here
	if headers != nil {
		logger.Info("Adding Headers")

		for _, currHeader := range headers {
			msg := "Adding Header " + currHeader.headerName + " with Value " + currHeader.headerValue
			logger.Info(msg)
			req.Header.Add(currHeader.headerName, currHeader.headerValue)
		}
	}

	logger.Info("Making Call")

	resp, err := client.Do(req)

	if err != nil {
		logger.Error(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Error(err.Error())
	}

	return resp.StatusCode, bodyBytes

}
