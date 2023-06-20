package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Chat struct {
	IsBot  bool
	Msg    string
	Source string
}

var chat []Chat
var chatMutex sync.Mutex

func handleInputChanges(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	question := strings.TrimSpace(string(body))
	if question == "" {
		http.Error(w, "Please enter valid input and try again.", http.StatusBadRequest)
		return
	}

	chatMutex.Lock()
	chat = append(chat, Chat{IsBot: false, Msg: question})
	chatMutex.Unlock()

	// Call your AI function here and get the response
	// For example: answer, source := callAI(question)

	// Add the AI response to the chat
	chatMutex.Lock()
	chat = append(chat, Chat{IsBot: true, Msg: answer, Source: source})
	chatMutex.Unlock()

	response := map[string]string{
		"answer": answer,
		"source": source,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/get_answer", handleInputChanges)
	fmt.Println("Server listening on port 5000")
	http.ListenAndServe(":5000", nil)
}
