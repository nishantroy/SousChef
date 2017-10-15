package mas

import (
	//"fmt"
	"net/http"
	"fmt"
)

func init() {

	//User Handlers
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)

	// Recipe Handlers
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)

	// Default
	http.HandleFunc("/", handler)
	//temp()

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err.Error())
}
