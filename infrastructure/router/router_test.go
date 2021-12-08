package router

import (
	"testing"

	"github.com/FernandoGal25/academy-go-q42021/infrastructure/mock"
	"github.com/FernandoGal25/academy-go-q42021/interface/controller"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	app := controller.AppController{
		Pokemon: mock.NewMockPokemonAction(mockCtrl),
	}

	want := echo.New()
	got := NewRouter(app)
	assert.IsType(t, want, got, "Test new router, expected: %v, got: %v", want, got)
}
