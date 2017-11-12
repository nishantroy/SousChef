package mas

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")
}

func init() {

	// USER HANDLERS

	// Weekly plan handlers
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/weekly_plan_create", handleCreateWeeklyPlan)
	http.HandleFunc("/api/v1/users/update_meal", handleUpdateMeal)

	// Shopping list handlers
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)
	http.HandleFunc("/api/v1/users/shopping_list_create", handleCreateShoppingList)
	http.HandleFunc("/api/v1/users/item_checked", handleCheckGroceryItem)
	http.HandleFunc("/api/v1/users/item_unchecked", handleUncheckGroceryItem)

	// User profile handlers
	http.HandleFunc("/api/v1/users/create_profile", handleCreateProfile)
	http.HandleFunc("/api/v1/users/get_profile", handleGetProfile)
	http.HandleFunc("/api/v1/users/update_profile", handleUpdateProfile)

	// Alexa handlers
	http.HandleFunc("/api/v1/alexa/get_alexa_auth_token", handleGenerateAlexaAuthToken)
	http.HandleFunc("/api/v1/alexa/authorize_alexa", handleAlexaAuth)
	http.HandleFunc("/api/v1/alexa/get_recipe_details", handleGetRecipeForAlexa)

	// User favorites handlers
	http.HandleFunc("/api/v1/users/add_favorite", handleAddFavoriteRecipe)
	http.HandleFunc("/api/v1/users/get_favorites", handleGetFavoriteRecipes)
	http.HandleFunc("/api/v1/users/delete_favorite", handleDeleteFavoriteRecipe)

	// User persistence handlers
	http.HandleFunc("/api/v1/users/save_current_recipe_progress", handleSaveCurrentRecipeProgress)
	http.HandleFunc("/api/v1/users/get_current_recipe_progress", handleGetCurrentRecipeProgress)
	http.HandleFunc("/api/v1/users/delete_current_recipe_progress", handleDeleteCurrentRecipeProgress)

	// Recipe handlers
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)
	http.HandleFunc("/api/v1/recipes/recipe_changes", handleGetRecipeChanges)

	// Default
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err.Error())
}
