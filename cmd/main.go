/*https://dev.to/karanpratapsingh/build-a-rest-api-with-go-for-beginners-3gp*/

package main

import (
	"log"
	"net/http"

	"github.com/RajaKAhmed/RegID/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/Registration", handlers.GetAllRegistrations).Methods(http.MethodGet)
	router.HandleFunc("/Registration{AppID}", handlers.GetRegistration).Methods(http.MethodGet)
	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
