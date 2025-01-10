package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type JsonResponse struct {
	Data interface{} `json:"data"`
}

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
}

func translateValidationErrors(err error) map[string]string {
	errs := make(map[string]string)
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, err := range validationErrors {
			fieldName := err.Field()
			switch err.Tag() {
			case "required":
				errs[fieldName] = fieldName + " is required"
			case "eth_addr":
				errs[fieldName] = fieldName + " is not a valid Ethereum address"
			default:
				errs[fieldName] = fieldName + " is invalid"
			}
		}
	}

	return errs
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	if err := writeJSON(w, status, &ErrorResponse{Error: message}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	return writeJSON(w, status, &JsonResponse{Data: data})
}
