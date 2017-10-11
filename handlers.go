package main

import (
	"net/http"
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"encoding/json"
)

const (
	fireURL string = "https://souschef-182502.firebaseio.com"
	authToken string = "4maH9UaAtODP5C64FUCpn51Y6kaKSjeSCIHuPZ5y"
)

var (
	f = firego.New(fireURL, nil)
)

type User struct {
	Name    string `json:"name"`
	Recipes interface{} `json:"recipes"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results
	user := &User{}

	f.Auth(authToken)
	
	if err := f.Child("users/" + userID).Value(user); err != nil {
		fmt.Print(w, err)
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