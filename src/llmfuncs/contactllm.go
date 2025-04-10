package llmfuncs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ollamaAPIURL = "http://localhost:11434/api/generate"

type RequestPayload struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ResponseData struct {
	Response string `json:"response"`
}

func ContactLLama(promptText string) (string, error) {
	payload := RequestPayload{
		Model:  "llama3.2",
		Prompt: promptText,
		Stream: false,
	}

	// Erstellen des JSON-Payloads
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("fehler beim Erstellen des JSON-Payloads: %w", err)
	}

	// Senden der Anfrage an die Llama-API
	resp, err := http.Post(ollamaAPIURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("fehler bei der Anfrage: %w", err)
	}
	defer resp.Body.Close()

	// Lesen der Antwort
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("fehler beim Lesen der Antwort: %w", err)
	}

	// Verarbeiten der Antwort
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("fehler: %d", resp.StatusCode)
	}

	// Extrahieren der Antwort aus dem JSON
	var responseData ResponseData
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("fehler beim Verarbeiten der JSON-Antwort: %w", err)
	}

	// Rückgabe der Antwort
	return responseData.Response, nil
}
