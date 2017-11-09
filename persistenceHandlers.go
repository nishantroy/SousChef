package mas

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func handleSaveCurrentRecipeProgress(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := SaveCurrentRecipeProgress(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func handleGetCurrentRecipeProgress(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	progress, err := GetCurrentRecipeProgress(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(progress)
}

func handleDeleteCurrentRecipeProgress(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := DeleteCurrentRecipeProgress(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}