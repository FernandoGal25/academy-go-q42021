package repository

import (
	"io"
	"strconv"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/datastore"
)

/*
	Repository that handles the operations of pokemon in a CSV file.
*/
type CSVPokemonRepository struct {
	Handler *datastore.CSVHandler
}

/*
	Returns an instance of CSVPokemonRepository.
*/
func NewCSVPokemonRepository(h *datastore.CSVHandler) CSVPokemonRepository {
	return CSVPokemonRepository{h}
}

func buildPokemon(p *model.Pokemon, data []string, ID uint64) error {
	p.ID = ID

	p.Name = data[1]

	h, err := strconv.Atoi(data[2])
	if err != nil {
		return customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID HEIGHT", Err: err}
	}
	p.Height = h

	w, err := strconv.Atoi(data[3])
	if err != nil {
		return customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID WEIGHT", Err: err}
	}
	p.Weight = w

	o, err := strconv.Atoi(data[4])
	if err != nil {
		return customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID ORDER", Err: err}
	}
	p.Order = o

	b, err := strconv.Atoi(data[5])
	if err != nil {
		return customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID BASE_EXPERIENCE", Err: err}
	}
	p.BaseExperience = b

	return nil
}

/*
	Searches the pokemon which ID belongs to.
*/
func (r CSVPokemonRepository) FindByID(ID uint64) (*model.Pokemon, error) {
	if err := r.Handler.BuildHandler(); err != nil {
		return nil, err
	}
	defer r.Handler.Close()
	var i model.Pokemon

	for {
		record, err := r.Handler.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, customErrors.CSVFormatError{Message: "CANNOT READ CSV", Err: err}
		}

		parsedID, err := strconv.ParseUint(record[0], 10, 16)
		if err != nil {
			return nil, customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID ID", Err: err}
		}
		if parsedID == ID {
			if err := buildPokemon(&i, record, ID); err != nil {
				return nil, err
			}
			break
		}
	}

	if i.ID == 0 {
		return nil, customErrors.EntityNotFoundError{
			Message: "THE SEARCHED POKEMON HAS NOT BEEN REGISTERED IN THIS DATABASE",
		}
	}

	return &i, nil
}

/*
	Returns all pokemon registered in the CSV file.
*/
func (r CSVPokemonRepository) FetchAll() ([]model.Pokemon, error) {
	err := r.Handler.BuildHandler()

	if err != nil {
		return nil, err
	}

	var collection = []model.Pokemon{}

	records, err := r.Handler.ReadAll()

	if err != nil {
		return nil, err
	}

	for _, record := range records {
		ID, err := strconv.ParseUint(record[0], 10, 16)
		if err != nil {
			return nil, customErrors.CSVFormatError{Message: "CSV ITEM DOES NOT POSSESS A VALID ID", Err: err}
		}
		var i model.Pokemon
		if err := buildPokemon(&i, record, ID); err != nil {
			return nil, err
		}

		collection = append(collection, i)
	}

	return collection, nil
}
