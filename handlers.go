package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

// HANDLERS FOR USER ENDPOINTS

// Takes in a user ID, fetches meal plan for the week from Firebase and returns
func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wp, err := getWeeklyPlanForUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wp.Days)
}

// Takes in a user ID, fetches profile, gets a weekly plan from the API, and saves to Firebase
func handleCreateWeeklyPlan(w http.ResponseWriter, req *http.Request) {

	if err := createWeeklyPlanForUser(req); err != nil {
		fmt.Fprintln(w, err)
	}

}

// Will take in a user ID, fetch his/her shopping list from Firebase and return
func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shopList, err := getShoppingListForUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shopList)
}

// Will take in comma separated list of recipe IDs chosen by user, fetch the ingredients,
// do unit conversions, create shopping list, and save to Firebase
func handleCreateShoppingList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := createShoppingListForUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintln(w, "SUCCESS!")
}

func handleCheckGroceryItem(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateGroceryItemDoneForUser(req, true)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleUncheckGroceryItem(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateGroceryItemDoneForUser(req, false)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

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

func handleUpdateMeal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateMeal(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "ERROR IN HELPER METHOD :", err)
		return
	}
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
