package service

import (
	"encoding/json"
	"net/http"
)

func GetJSONSimpleStringMessage(w http.ResponseWriter, m string) {
	// Create a data structure to be returned as JSON
	data := struct {
		Message string `json:"message"`
	}{
		Message: m,
	}

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	setJSONHeadersAndWrite(w, jsonData)
}

func GetJSONData[T any](w http.ResponseWriter, data T) {
	// Marshal the data into JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	setJSONHeadersAndWrite(w, jsonData)
}

func setJSONHeadersAndWrite(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
