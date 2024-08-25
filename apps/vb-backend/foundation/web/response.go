package web

import (
	"encoding/json"
	"net/http"
)

type responseStatus string

const (
	statusOk    responseStatus = "OK"
	statusError responseStatus = "ERROR"
)

type responseSuccess[T any] struct {
	Data   T              `json:"data"`
	Status responseStatus `json:"status"`
}

type responseError[T any] struct {
	Error  T              `json:"error"`
	Status responseStatus `json:"status"`
}

func Error[T any](data T) responseSuccess[T] {
	return responseSuccess[T]{Data: data, Status: statusOk}
}

func Success[T any](err T) responseError[T] {
	return responseError[T]{Error: err, Status: statusError}
}

// Respond converts a Go value to JSON and sends it to the client.
func Respond[T any](w http.ResponseWriter, data responseSuccess[T], statusCode int) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
