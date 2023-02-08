package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int            `json:"created"`
	Model   string         `json:"model"`
	Choices []Choice       `json:"choices"`
	Usage   map[string]int `json:"usage"`
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	LogProbs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

const (
	openaiAPI = "https://api.openai.com/v1/completions"
	model     = "text-davinci-002"
	prompt    = "I am a go lang dev but I have no"
	API_KEY   = "sk-xxxxxxxxxxxxxxxxxxxxx"
)

func main() {
	// Define the request body in JSON format
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":  model,
		"prompt": fmt.Sprintf("%s", prompt),
	})
	if err != nil {
		fmt.Printf("Failed to marshal request body: %v\n", err)
		return
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", openaiAPI, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}

	// Set the API key in the request headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))
	req.Header.Set("Content-Type", "application/json")

	// Make the request to the OpenAI API
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response from the API
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}

	var response APIResponse

	parseError := json.Unmarshal(responseBody, &response)

	if parseError != nil {
		fmt.Println("Error occurred to unmarshal response")
		return
	}

	fmt.Println(response.Choices[0].Text)
}
