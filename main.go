package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const openAIKey = "YOUR_OPENAI_API_KEY" // Replace with your API key

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ask", askHandler)
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	response, err := queryOpenAI(string(body))
	if err != nil {
		http.Error(w, "Error querying OpenAI", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, response)
}

func queryOpenAI(query string) (string, error) {
	type Payload struct {
		Prompt    string `json:"prompt"`
		MaxTokens int    `json:"max_tokens"`
	}

	data := Payload{
		Prompt:    query,
		MaxTokens: 150,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci/completions", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-OK HTTP status: %d", resp.StatusCode)
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return "", err
		}
		log.Printf("Response body: %s", string(bodyBytes))
		return "", fmt.Errorf("Non-OK HTTP status: %d", resp.StatusCode)
	}

	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return "", err
	}

	response, ok := respData["choices"].([]interface{})
	if !ok || len(response) == 0 {
		return "", fmt.Errorf("invalid response from OpenAI")
	}

	choice, ok := response[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid choice format")
	}

	text, ok := choice["text"].(string)
	if !ok {
		return "", fmt.Errorf("invalid text format")
	}

	return text, nil
}
