package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error string `json:"error,omitempty"`
}

const (
	StatusOK = "ok"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errorMessages []string // Slice

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, err.Field() + " is required")
		case "min":
			// errorMessages = append(errorMessages, err.Field() + " must be at least " + err.Param())
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
		case "max":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param()))

		default:
			// errorMessages = append(errorMessages, err.Error())
			// errorMessages = append(errorMessages, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.ActualTag()))
			errorMessages = append(errorMessages, fmt.Sprintf("Field validation for '%s' ib invalid", err.Field()))
		}
	}


	return Response{
		Status: StatusError,
		Error: strings.Join(errorMessages, ", "),
	}
}
