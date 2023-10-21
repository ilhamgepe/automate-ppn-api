package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
)

func SendHttpRequest(url string, method HTTPMethod, body []byte) (*http.Response, time.Duration, error) {
	client := &http.Client{}
	var req *http.Request
	var err error

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("API_KEY")),
	}

	if body != nil {
		req, err = http.NewRequest(string(method), url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(string(method), url, nil)
	}
	if err != nil {
		return nil, 0, err
	}

	//tambahkan header yang ada
	for key, value := range headers {

		req.Header.Add(key, value)
	}

	//waktu mulai
	startTime := time.Now()

	// hit request
	response, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	//waktu selesai
	endTime := time.Now()

	responseTime := endTime.Sub(startTime)

	return response, responseTime, nil
}
