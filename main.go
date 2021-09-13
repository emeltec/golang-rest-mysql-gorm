package main

import (
	"log"
	"net/http"

	"github.com/emeltec/golang-rest-mysql-gorm/commons"
	"github.com/emeltec/golang-rest-mysql-gorm/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Server running on port 9000")
	log.Println(server.ListenAndServe())
}
