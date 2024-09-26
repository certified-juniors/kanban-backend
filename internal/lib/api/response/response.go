package response

import (
	"encoding/json"
	"net/http"
)

// Response структура ответа
type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

const (
	StatusOK        = "OK"
	StatusNoContent = "No Content"
	StatusError     = "Error"
)

// WriteJSONResponse принимает статус ответа, данные, сообщение об ошибке и ошибки валидации и записывает JSON-ответ.
func WriteJSONResponse(w http.ResponseWriter, status int, errMsg string, validationErrors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var response Response

	switch status {
	case http.StatusOK:
		response.Status = StatusOK
	case http.StatusNoContent:
		response.Status = StatusNoContent
	case http.StatusCreated:
		response.Status = StatusOK
	default:
		response.Status = StatusError
	}

	if errMsg != "" || validationErrors != nil {
		response.Error = errMsg
		response.Errors = validationErrors
	}

	json.NewEncoder(w).Encode(response)
}

// WriteJSON записывает Response в формате JSON в http.ResponseWriter
func WriteJSON(w http.ResponseWriter, statusCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
