package contracts

import "github.com/FernandoGal25/academy-go-q42021/domain/model"

// PokemonRepository is an interface of pokemon repository, defines the
// methods of the repository needed to handle the entity in question.
type PokemonRepository interface {
	FindByID(id int) (*model.Pokemon, error)
	FetchAll() ([]model.Pokemon, error)
	Persist(*model.Pokemon) error
	FetchConcurrently(f map[string]interface{}) ([]model.Pokemon, error)
}
