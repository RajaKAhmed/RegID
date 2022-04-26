package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"math/rand"
	"github.com/RajaKAhmed/RegID/pkg/mocks"
)

func RandomString(n int) string {GetAllRegistrations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Registration)
}
