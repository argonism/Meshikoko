package main

import (
	"app/config"
	"app/database"
	"app/models"
	"app/server"
	"flag"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(true, models.GetModels()...)
	if err := server.Init(); err != nil {
		panic(err)
	}
}
