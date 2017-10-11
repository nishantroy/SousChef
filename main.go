package mas

import (
	"fmt"
	"net/http"
)



func init() {
	http.HandleFunc("/api/v1/users/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/api/v1/users/shopping_list", handleGetShoppingList)
	http.HandleFunc("/api/v1/recipes/recipe_steps", handleGetRecipeSteps)
	http.HandleFunc("/api/v1/recipes/recipe_details", handleGetRecipeDetails)
	http.HandleFunc("/", handler)


	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}