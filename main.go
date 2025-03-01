package main

import (
	"log"
	"net/http"

	"github.com/faizkhan-06/gozap/config"
	"github.com/faizkhan-06/gozap/src/routes"
)

func main() {
	config.DbConnect()
	router := routes.RegisterRoutes()
	
	log.Println("Server listning on 3000")
	if err := http.ListenAndServe(":3000", router); err != nil{
		log.Fatal("Server Failed")
	}

}