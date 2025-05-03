package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"bufio"
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
		"stream": true,
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

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}

		var chunk map[string]interface{}
		if err := json.Unmarshal(line, &chunk); err != nil {
			continue
		}

		if respText, ok := chunk["response"].(string); ok {
			w.Write([]byte(respText))
			flusher.Flush()
		}

		if done, ok := chunk["done"].(bool); ok && done {
			break
		}
	}
}
