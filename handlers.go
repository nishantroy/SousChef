package mas

import (
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"

	"encoding/json"

	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine"
)

const (
	fireURL string = "https://souschef-182502.firebaseio.com"
	authToken string = "4maH9UaAtODP5C64FUCpn51Y6kaKSjeSCIHuPZ5y"
)

type User struct {
	Name    string      `json:"name"`
	Recipes interface{} `json:"recipes"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	// Make call to API or to Database here, and then write out results
	var user interface{}

	f.Auth(authToken)

	if err := f.Child("users/" + userID).Value(&user); err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(user)
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
