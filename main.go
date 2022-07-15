package main

import (
	"github.com/Wmartevi/web-api-go/database"
	"github.com/Wmartevi/web-api-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
