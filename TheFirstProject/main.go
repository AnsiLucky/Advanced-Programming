package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequestHome)

	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server Error:", err)
	}
}

func handleRequestHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			sendResponse(w, "400")
			return
		}
		if message.Message == "" {
			sendResponse(w, "400")
			return
		}
		fmt.Println(message.Message)
		sendResponse(w, "200")
	} else {
		http.Error(w, "Этот метод не обрабатывается", http.StatusMethodNotAllowed)
	}
}

func sendResponse(w http.ResponseWriter, status string) {
	var message string
	if status == "200" {
		message = "Данные успешно приняты"
		w.WriteHeader(http.StatusOK)
	} else {
		message = "Некорректное JSON-сообщение"
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{status, message})
}
