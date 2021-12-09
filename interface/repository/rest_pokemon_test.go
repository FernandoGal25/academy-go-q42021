package repository_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/domain/model"
	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/mock"
	"github.com/FernandoGal25/academy-go-q42021/interface/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//Mock io.Reader in order to generate ioutil.ReadAll to make an error.
type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, assert.AnError
}
func TestRESTPokemonRepository_FindByID(t *testing.T) {
	pokemon := model.Pokemon{
		ID:             1,
		Name:           "bulbasaur",
		Height:         7,
		Weight:         69,
		Order:          1,
		BaseExperience: 64,
	}
	endpoint := "/pokemon/1"
	badJson := []byte("Bad JSON marshal")
	errBadJson := json.Unmarshal(badJson, &pokemon)
	tests := []struct {
		name       string
		key        int
		prepare    func(client *mock.MockHTTPClient)
		wantResult *model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			key:  1,
			prepare: func(client *mock.MockHTTPClient) {
				body, _ := json.Marshal(pokemon)
				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(bytes.NewReader(body)),
				}, nil)
			},
			wantResult: &pokemon,
			wantErr:    nil,
		},
		{
			name: "Bad JSON test",
			key:  1,
			prepare: func(client *mock.MockHTTPClient) {
				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(bytes.NewReader(badJson)),
				}, nil)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed to parse json body from HTTP response: %v", endpoint), Err: errBadJson},
		},
		{
			name: "Bad Response Body",
			key:  1,
			prepare: func(client *mock.MockHTTPClient) {

				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(errReader(0)),
				}, nil)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Cannot read HTTP response body from %v", endpoint), Err: assert.AnError},
		},
		{
			name: "Bad HTTP Response",
			key:  1,
			prepare: func(client *mock.MockHTTPClient) {

				client.EXPECT().Get(endpoint).Return(nil, assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed request at %v", endpoint), Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		client := mock.NewMockHTTPClient(mockCtrl)
		repo := repository.NewRestPokemonRepository(client)
		if tt.prepare != nil {
			tt.prepare(client)
		}
		got, err := repo.FindByID(1)
		assert.Equal(t, tt.wantResult, got, "%v , expected: %v, got: %v", tt.name, tt.wantResult, got)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}

func TestRESTPokemonRepository_FetchAll(t *testing.T) {
	pokemons := []model.Pokemon{{
		ID:             1,
		Name:           "bulbasaur",
		Height:         7,
		Weight:         69,
		Order:          1,
		BaseExperience: 64,
	}}
	endpoint := "/pokemon"
	badJson := []byte("Bad JSON marshal")
	errBadJson := json.Unmarshal(badJson, &pokemons)
	tests := []struct {
		name       string
		prepare    func(client *mock.MockHTTPClient)
		wantResult []model.Pokemon
		wantErr    error
	}{
		{
			name: "Correct test",
			prepare: func(client *mock.MockHTTPClient) {
				body, _ := json.Marshal(&pokemons)
				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(bytes.NewReader(body)),
				}, nil)
			},
			wantResult: pokemons,
			wantErr:    nil,
		},
		{
			name: "Bad JSON test",
			prepare: func(client *mock.MockHTTPClient) {
				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(bytes.NewReader(badJson)),
				}, nil)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed to parse json body from HTTP response: %v", endpoint), Err: errBadJson},
		},
		{
			name: "Bad Response Body",
			prepare: func(client *mock.MockHTTPClient) {

				client.EXPECT().Get(endpoint).Return(&http.Response{
					Body: ioutil.NopCloser(errReader(0)),
				}, nil)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Cannot read HTTP response body from %v", endpoint), Err: assert.AnError},
		},
		{
			name: "Bad HTTP Response",
			prepare: func(client *mock.MockHTTPClient) {

				client.EXPECT().Get(endpoint).Return(nil, assert.AnError)
			},
			wantResult: nil,
			wantErr:    customErrors.ErrHTTPRequest{Message: fmt.Sprintf("Failed request at %v", endpoint), Err: assert.AnError},
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, tt := range tests {
		client := mock.NewMockHTTPClient(mockCtrl)
		repo := repository.NewRestPokemonRepository(client)
		if tt.prepare != nil {
			tt.prepare(client)
		}
		got, err := repo.FetchAll()
		assert.Equal(t, tt.wantResult, got, "%v , expected: %v, got: %v", tt.name, tt.wantResult, got)
		assert.Equal(t, tt.wantErr, err, "%v , expected: %v, got: %v", tt.name, tt.wantErr, err)
	}
}
