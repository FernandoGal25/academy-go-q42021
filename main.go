package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"academy_bootcamp/config"
	"academy_bootcamp/config/registry"
	"academy_bootcamp/infrastructure/datastore"
	"academy_bootcamp/infrastructure/router"
)

func main() {
	config.ReadConfig()

	h := datastore.NewCSVHandler(config.C.CSV.Path)

	r := registry.NewRegistry(h)

	e := echo.New()
	e = router.NewRouter(e, r.Register())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Port)
	if err := e.Start(":" + config.C.Server.Port); err != nil {
		log.Fatalln(err)
	}
}
