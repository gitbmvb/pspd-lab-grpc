package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AskHandler(w http.ResponseWriter, r *http.Request) {
	type InputPayload struct {
		Input string `json:"input"`
	}

	var payload InputPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reqBody := map[string]interface{}{
		"model":  "llama3.2",
		"prompt": payload.Input,
		"stream": false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	url := "http://llm_server:11434/api/generate"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error contacting model:", err)
		http.Error(w, "Failed to contact model", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("Model returned error:", string(bodyBytes))
		http.Error(w, "Model returned error", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read model response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
