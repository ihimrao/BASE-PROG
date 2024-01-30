package middlewares

import (
	"encoding/json"
	"net/http"
)

func AuthorizationResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 401, Message: msg}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(writer).Encode(temp)
}

func SuccessResponse(msg string, w http.ResponseWriter) {
	type error struct {
		StatusCode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &error{StatusCode: http.StatusOK, Message: msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}

func ErrorResponse(msg string, w http.ResponseWriter) {
	type error struct {
		StatusCode int    `json:"status"`
		Message    string `json:"msg"`
	}

	temp := &error{StatusCode: http.StatusBadRequest, Message: msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(temp)

}
