package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SetHTTPStatus(w http.ResponseWriter, status int, message string, data int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	jsonResp, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("json marshal err: %s", err)
	}
	w.Write(jsonResp)
}

func SetHTTPResponse(w http.ResponseWriter, status int, data *[]map[string]int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json marshal err: %s", err)
	}
	w.Write(jsonResp)
}
