package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/bjayanta/students-api/internal/types"
	"github.com/bjayanta/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		slog.Info("New student request")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			// response.WriteJson(w, http.StatusBadGateway, response.GeneralError(err))
			response.WriteJson(w, http.StatusBadGateway, response.GeneralError(fmt.Errorf("request body missing")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // Type cast
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}


		// w.Write([]byte("Welcome to students api!"))
		// response.WriteJson(w, http.StatusCreated, student)
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}