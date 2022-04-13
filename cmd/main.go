package main

import (
	"log"
	"net/http"

	"github.com/RajaKAhmed/RegID/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/Registration", handlers.GetRegistration).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
