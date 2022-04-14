package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/RajaKAhmed/RegID/pkg/mocks"
	"github.com/RajaKAhmed/RegID/pkg/models"
)

func AddRegistration(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var registration models.Registration
	json.Unmarshal(body, &registration)

	// Append to the Book mocks
	registration.AppID = rand.Intn(100)
	mocks.Registration = append(mocks.Registration, registration)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
