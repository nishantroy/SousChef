package main

import (
	//"fmt"
	"net/http"
)

func main() {

	//User Handlers
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)

	// Recipe Handlers
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)

	// Default
	http.HandleFunc("/", temp)
	//temp()
}
