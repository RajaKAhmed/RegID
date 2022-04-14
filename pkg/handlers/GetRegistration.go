package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RajaKAhmed/RegID/pkg/mocks"
)

func GetRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Registration)
}
