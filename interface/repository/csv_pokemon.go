package repository

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/FernandoGal25/academy-go-q42021/helpers"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/datastore"
)

// Repository that handles the operations of pokemon in a CSV file.
type CSVPokemonRepository struct {
	Handler *datastore.CSVHandler
}

// Returns an instance of CSVPokemonRepository.
func NewCSVPokemonRepository(h *datastore.CSVHandler) CSVPokemonRepository {
	return CSVPokemonRepository{h}
}

func (r CSVPokemonRepository) initializeHandler() error {
	if err := r.Handler.BuildHandler(); err != nil {
		return customErrors.ErrDatastoreWrapper{Message: "Could not initialize CSVHandler", Err: err}
	}

	return nil
}

// NOTE: This function could be a generic builder of structs.
func (r CSVPokemonRepository) buildPokemon(p *model.Pokemon, data []string, ID int) error {
	p.ID = ID

	rv := reflect.ValueOf(p).Elem()

	for key, sField := range r.Handler.Schema {
		if sField == "id" {
			continue
		}

		cField := helpers.SnakeCaseToCamelCase(sField)
		vv := rv.FieldByName(cField)

		if !vv.IsValid() {
			return customErrors.ErrCSVFormat{Message: fmt.Sprintf("Passed struct does not possess field `%v` defined as `%v` in CSV", cField, sField)}
		}

		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(data[key], 10, 64)
			if err != nil {
				return customErrors.ErrCSVFormat{
					Message: fmt.Sprintf("CSV item does not possess a valid %v", cField),
					Err:     err,
				}
			}
			vv.SetInt(v)
		case reflect.String:
			vv.SetString(data[key])
		}
	}

	return nil
}

func unBuildPokemon(p *model.Pokemon) []string {
	return []string{
		strconv.Itoa(p.ID),
		p.Name,
		strconv.Itoa(p.Height),
		strconv.Itoa(p.Weight),
		strconv.Itoa(p.Order),
		strconv.Itoa(p.BaseExperience),
	}
}

// Searches the pokemon which ID belongs to.
func (r CSVPokemonRepository) FindByID(ID int) (*model.Pokemon, error) {
	if err := r.initializeHandler(); err != nil {
		return nil, err
	}
	defer r.Handler.Close()
	var p model.Pokemon

	for {
		record, err := r.Handler.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, customErrors.ErrDatastoreWrapper{Message: "Failed to read from datastore", Err: err}
		}

		parsedID, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, customErrors.ErrCSVFormat{Message: "CSV item does not possess a valid ID", Err: err}
		}
		if parsedID == ID {
			if err := r.buildPokemon(&p, record, ID); err != nil {
				return nil, err
			}
			break
		}
	}

	if p.ID == 0 {
		return nil, customErrors.ErrEntityNotFound{
			Message: "The searched pokemon is still not registered in the pokedex",
			Err:     nil,
		}
	}

	return &p, nil
}

// Returns all pokemon registered in the CSV file.
func (r CSVPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	if err := r.initializeHandler(); err != nil {
		return nil, err
	}
	defer r.Handler.Close()
	var collection = []model.Pokemon{}

	records, err := r.Handler.ReadAll()

	if err != nil {
		return nil, customErrors.ErrDatastoreWrapper{Message: "Failed to fetch all items", Err: err}
	}

	for _, record := range records {
		ID, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, customErrors.ErrCSVFormat{Message: "CSV item does not possess a valid ID", Err: err}
		}
		var i model.Pokemon
		if err := r.buildPokemon(&i, record, ID); err != nil {
			return nil, err
		}

		collection = append(collection, i)
	}

	return collection, nil
}

// Saves pokemon on csv.
func (r CSVPokemonRepository) Persist(p *model.Pokemon) error {
	if err := r.initializeHandler(); err != nil {
		return err
	}
	defer r.Handler.Close()

	if err := r.Handler.Write(unBuildPokemon(p)); err != nil {
		return customErrors.ErrDatastoreWrapper{Message: "Failed to persist item", Err: err}
	}

	return nil
}
