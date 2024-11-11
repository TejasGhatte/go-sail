package scripts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TejasGhatte/go-sail/internal/prompts"
	"github.com/TejasGhatte/go-sail/internal/signals"
)

func Signup(ctx context.Context) error {
	ctx = signals.HandleCancellation(ctx)

	username, email, password, err := prompts.PromptUserSignupDetails(ctx)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
		"plan": "Premium",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	response, err := sendSignupReq("url", jsonData)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var responseBody struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	if response.StatusCode != 200 {
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&responseBody); err == nil {
			return err
		} else {
			fmt.Printf("Error decoding response: %v\n", err)
		}
		return fmt.Errorf("error calling mailer")
	}
	fmt.Println("Signup successful!")
	return nil
}

func sendSignupReq(URL string, data []byte) (*http.Response, error) {

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