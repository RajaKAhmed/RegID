package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RajaKAhmed/RegID/pkg/mocks"
)

func GetRegistration(w http.ResponseWriter, r *http.Request) {
	// Read Registration ID
	// vars := mux.Vars(r)
	// AppID, _ := strconv.Atoi(vars["AppID"])
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
