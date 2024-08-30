package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Prompt   string
	Message  string
	TimeSent int64
	TimeRcvd int64
	Source   string
}

const LLM_API = "https://api-inference.huggingface.co/models/google/gemma-7b"

// server runs on port 4000
func StartServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 4000...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt query parameter is missing", http.StatusBadRequest)
		return
	}

	timeSent := time.Now().Unix()
	message, err := fetchLLM(prompt)
	if err != nil {
		http.Error(w, "Failed to fetch result", http.StatusInternalServerError)
		return
	}

	timeRcvd := time.Now().Unix()
	response := Response{
		Prompt:   prompt,
		Message:  message,
		TimeSent: timeSent,
		TimeRcvd: timeRcvd,
		Source:   "gemma-7b",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func fetchLLM(prompt string) (string, error) {
	token := os.Getenv("HUGGINGFACE_API_TOKEN")
	if token == "" {
		return "", fmt.Errorf("API token not provided")
	}

	requestBody := map[string]string{"inputs": prompt}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", LLM_API, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", response.StatusCode, body)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(responseBody), nil
}
