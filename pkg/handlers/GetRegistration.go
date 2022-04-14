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
	vars := mux.Vars(r)
	AppID, _ := strconv.Atoi(vars["AppID"])
	// Iterate over all Registrations
	for _, register := range mocks.Registration {
		if register.AppID == AppID {
			// If registration ID is the same send Info as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(register)
			break
		}
	}
}
