package mas

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the Sous Chef API!")
}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello User %s! Your weekly plan is: _____", userID)
}

func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello User %s! Your shopping list is: _____", userID)
}

func handleGetRecipeSteps(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello! The steps for recipe %s are: _____", recipeID)
}

func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello! The details for recipe %s are: _____", recipeID)
}