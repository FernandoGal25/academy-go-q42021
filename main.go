package main

import (
	"log"

	"github.com/FernandoGal25/academy-go-q42021/config"
	"github.com/FernandoGal25/academy-go-q42021/config/registry"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/datastore"
	"github.com/FernandoGal25/academy-go-q42021/infrastructure/router"
)

func main() {
	c, err := config.ReadConfig()

	if err != nil {
		log.Fatalln(err)
		return
	}

	handler := datastore.NewCSVHandler(c.CSV.Path)
	client := datastore.NewHTTPClient(c.Rest.Api)
	r := registry.NewRegistry(handler, client)
	e := router.NewRouter(r.Register())

	if err := e.Start(":" + c.Server.Port); err != nil {
		log.Fatalln(err)
	}
}
