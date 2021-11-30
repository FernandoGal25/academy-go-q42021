package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
)

/*
	Wrapper of controllers
*/
type AppController struct {
	Pokemon PokemonController
}

func parseIDParam(c Context) (uint64, error) {
	key, err := strconv.ParseUint(c.Param("ID"), 10, 16)
	if err != nil {
		return 0, customErrors.InvalidRequestError{Message: "INVALID ID PARAM, MUST BE A NUMBER", Err: err}
	}

	return key, nil
}

func responseJSON(c Context, result interface{}, code int) error {
	if err := c.JSON(code, result); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func responseErrorJSON(c Context, err error) error {
	var status int
	var body interface{}

	switch err.(type) {
	case customErrors.InvalidRequestError:
		status = http.StatusBadRequest
		body = map[string]string{"ErrorMessage": err.Error()}
		if err := errors.Unwrap(err); err != nil {
			log.Fatalln(err)
		}
	case customErrors.DomainValidationError:
		status = http.StatusBadRequest
		body = map[string]string{"ErrorMessage": err.Error()}
		if err := errors.Unwrap(err); err != nil {
			log.Fatalln(err)
		}
	case customErrors.EntityNotFoundError:
		status = http.StatusNotFound
		body = map[string]string{"ErrorMessage": err.Error()}
		if err := errors.Unwrap(err); err != nil {
			log.Fatalln(err)
		}
	case customErrors.CSVFormatError:
		status = http.StatusInternalServerError
		body = map[string]string{"ErrorMessage": err.Error()}
		if err := errors.Unwrap(err); err != nil {
			log.Fatalln(err)
		}
	default:
		status = http.StatusInternalServerError
		body = map[string]string{"ErrorMessage": customErrors.DEFAULT_MESSAGE}
		log.Fatalln(err)
	}

	return responseJSON(c, body, status)
}
