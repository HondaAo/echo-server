package main

import (
	"log"

	database "github.com/HondaAo/snippet/src/db"
	envfile "github.com/HondaAo/snippet/src/env"
	"github.com/HondaAo/snippet/src/server"
	"github.com/labstack/echo/v4"
)

func main() {
	env := envfile.SetEnv()
	db := database.InitDB(env)
	e := echo.New()

	s := server.NewServer(e, db)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
