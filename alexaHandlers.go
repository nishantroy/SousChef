package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleGenerateAlexaAuthToken(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := GenerateAlexaAuthToken(req)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func handleAlexaAuth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := AlexaAuthorize(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(false)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userID)
}

func handleCheckAlexaAuth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := CheckAlexaAuth(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(false)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userID)
}

func handleGetRecipeForAlexa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rec, err := GetRecipeForAlexa(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("error occurred")
		return
	}

	json.NewEncoder(w).Encode(rec)
}
