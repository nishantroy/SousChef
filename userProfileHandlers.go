package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleCreateProfile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := createUserProfile(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func handleGetProfile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	u, err := getUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)

}

func handleUpdateProfile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateUserProfile(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func handleGenerateAlexaAuthToken(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := GenerateAlexaAuthToken(req)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func handleAlexaAuth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := AlexaAuthorize(req)

	w.WriteHeader(http.StatusOK)

	if err != nil {
		fmt.Fprintln(w, err)
	}
}
