package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RajaKAhmed/RegID/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetRegistration(w http.ResponseWriter, r *http.Request) {
	// Read Registration ID
	readID := mux.Vars(r)
	AppID, _ := strconv.Atoi(readID["AppID"])
	// Iterate over all Registrations
	for _, Registration := range mocks.Registration {
		if Registration.AppID == AppID {
			// If registration ID is the same send Info as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Registration)
			break
		}
	}
}
