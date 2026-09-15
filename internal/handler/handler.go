package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kvitonity/go-calc-service/internal/calculator"
)

type Response struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	op1Str := query.Get("op1")
	operator := query.Get("operator")
	op2Str := query.Get("op2")

	op1, err := strconv.ParseFloat(op1Str, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid op1")
		return
	}

	op2, err := strconv.ParseFloat(op2Str, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid op2")
		return
	}

	result, err := calculator.Calculate(op1, operator, op2)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Result: result}); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: message}); err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}
