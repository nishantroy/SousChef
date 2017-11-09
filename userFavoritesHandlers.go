package mas

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func handleAddFavoriteRecipe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := AddFavoriteRecipe(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func handleGetFavoriteRecipes(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	favorites, err := GetFavoriteRecipes(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favorites)
}

func handleDeleteFavoriteRecipe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := DeleteFavoriteRecipe(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}
