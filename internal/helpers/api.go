package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeReq(URL string, data []byte) (*http.Response, error) {
	apiToken := GetApiToken()

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-token", apiToken)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Generalized request function
func MakeRequest(action string, extraData map[string]interface{}) (*Response, error) {
	payload := map[string]interface{}{
		"action": action,
	}
	for key, value := range extraData {
		payload[key] = value
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %w", err)
	}

	response, err := MakeReq("url", jsonData)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer response.Body.Close()

	var responseBody Response
	if response.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&responseBody); err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return nil, fmt.Errorf("server error: %s", responseBody.Message)
	}

	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}

	return &responseBody, nil
}

func GetApiToken() string {
	username, err := GetKey("username")
	if err != nil {
		return ""
	}
	if username == "" {
		return ""
	}
	api_token, _ := GetKey(fmt.Sprintf("api-token-%s", username))
	return api_token
}