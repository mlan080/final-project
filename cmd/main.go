package main

import (
	"github.com/mlan080/final-project/internal/airports"
	"github.com/mlan080/final-project/internal/database"
	"github.com/mlan080/final-project/internal/server"
)

func main() {
	//needs to run first or will get an empty datastore
	db := database.New()
	airports.ParseCSV(db)
	server.Server(db)
}
