package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorSchema struct {
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message error) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)

	if message == nil {
		write(w, nil)
		return
	}

	p := &ErrorSchema{Message: message.Error()}
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	if string(data) == "null" {
		return
	}

	write(w, data)
}

func write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}
