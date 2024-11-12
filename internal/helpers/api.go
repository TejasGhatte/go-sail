package helpers

import (
	"bytes"
	"net/http"
)

func MakeReq(URL string, data []byte) (*http.Response, error) {

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-token", "something")

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}