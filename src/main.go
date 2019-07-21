package main

import (	
	"log"
	"booksapi/src/router"
	"booksapi/src/config"
)

func main() {
	r := router.Route()
	port := config.All["port"]	
	log.Println("API listening at http://localhost:" + port)
	r.Run(":" + port)
}