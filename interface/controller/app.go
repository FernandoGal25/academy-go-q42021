package controller

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
)

// Wrapper of controllers.
type AppController struct {
	Pokemon PokemonAction
}

type ErrorResponse struct {
	Message   string
	ErrorType string
}

func parseIDParam(c Context) (int, error) {
	key, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, customErrors.ErrInvalidRequest{Message: "Invalid ID param, must be a number", Err: err}
	}

	return key, nil
}

func responseJSON(c Context, result interface{}, code int) error {
	if err := c.JSON(code, result); err != nil {
		return err
	}

	return nil
}

func responseErrorJSON(c Context, err error) error {
	var status int
	var r []ErrorResponse

BuildResponse:
	for {
		switch err.(type) {
		case nil:
			break BuildResponse
		case customErrors.ErrInvalidRequest, customErrors.ErrDomainValidation:
			status = http.StatusBadRequest
		case customErrors.ErrEntityNotFound:
			status = http.StatusNotFound
		case customErrors.ErrHTTPRequest, customErrors.ErrCSVFormat:
			status = http.StatusInternalServerError
		}

		r = append(r, ErrorResponse{Message: err.Error(), ErrorType: reflect.TypeOf(err).String()})
		err = errors.Unwrap(err)
	}

	if status == 0 {
		status = http.StatusInternalServerError
	}

	return responseJSON(c, map[string][]ErrorResponse{"errors": r}, status)
}
