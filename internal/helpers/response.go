package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type apiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JsonResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	if status < 100 || status > 500 {
		w.WriteHeader(500)
		log.Fatalf("invalid response status: %d", status)
		return
	}

	if status == 500 {
		message = "Internal Server Error"
	}

	response := &apiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Printf("fatal error marshal json data")
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
