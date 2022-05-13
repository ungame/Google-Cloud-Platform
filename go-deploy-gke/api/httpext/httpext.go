package httpext

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Port int

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(HEADER_CONTENT_TYPE, MIME_JSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("unable to encode json: %s\n", err.Error())
	}
}

type JsonError struct {
	Error string `json:"error"`
}

func NewJsonError(err error) *JsonError {
	if err == nil {
		return nil
	}
	return &JsonError{
		Error: err.Error(),
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, NewJsonError(err))
}
