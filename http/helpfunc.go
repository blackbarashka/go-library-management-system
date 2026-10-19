package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func parseBoolPtr(s string) (*bool, error) {
	if s == "" {
		return nil, nil // фильтр не задан
	}
	v, err := strconv.ParseBool(s) // принимает "1","0","t","f","true","false", ...
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func sendError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errDro := ErrorResponse{
		Message: err.Error(),
		Time:    time.Now(),
	}
	errd := json.NewEncoder(w).Encode(errDro)
	if errd != nil {
		panic(errd)
	}

}