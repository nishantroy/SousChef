package mas

import (
	"net/http"
	"fmt"
)

func init() {

	//User Handlers

	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/weekly_plan_create", handleCreateWeeklyPlan)

	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)
	http.HandleFunc("/api/v1/users/shopping_list_create", handleCreateShoppingList)

	// Recipe Handlers
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)

	// Default
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err.Error())
}