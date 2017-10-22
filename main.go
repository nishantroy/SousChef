package mas

import (
	"fmt"
	"net/http"
)

func init() {

	// USER HANDLERS

	// Weekly plan handlers
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/weekly_plan_create", handleCreateWeeklyPlan)

	// Shopping list handlers
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)
	http.HandleFunc("/api/v1/users/shopping_list_create", handleCreateShoppingList)

	// User profile handlers
	http.HandleFunc("/api/v1/users/create_profile", handleCreateProfile)
	http.HandleFunc("/api/v1/users/update_profile", handleUpdateProfile)

	// RECIPE HANDLERS
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)

	// Default
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err.Error())
}
