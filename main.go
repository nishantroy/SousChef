package mas

import (
	"fmt"
	"net/http"
)

func init() {

	// User Handlers
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)

	// Recipe Handlers
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)

	// Default
	http.HandleFunc("/", handler)

	// Start routing
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err.Error())
}
