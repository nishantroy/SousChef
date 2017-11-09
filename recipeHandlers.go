package mas

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// HANDLERS FOR RECIPE ENDPOINTS
func handleGetRecipeChanges(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recipes, err := getRecipeChanges(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "ERROR IN HELPER METHOD :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)

}

// Takes in a recipeID, gets the instructions for it from the API/cache, and returns
func handleGetRecipeSteps(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recipeID := req.URL.Query().Get("recipe_id")

	recipe, err := getRecipeDetails(req, recipeID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe.Instructions)
}

// Takes in a recipeID, gets the details for it from the API/cache, and returns
func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")
	w.Header().Set("Content-Type", "application/json")

	recipe, err := getRecipeDetails(req, recipeID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)
}
