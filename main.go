package main

import (
	"github.com/SillasReis/go-rest-api/database"
	"github.com/SillasReis/go-rest-api/routes"
)

func main() {
	database.PsqlConnect()
	routes.HandleRequest()
}
