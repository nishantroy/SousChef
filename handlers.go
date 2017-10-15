package main

import (
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"

	"encoding/json"

	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine"
	"os"
)

var (
	fireURL   string = os.Getenv("FIREBASE_URL")
	authToken string = os.Getenv("FIREBASE_AUTH_TOKEN")
)

type User struct {
	Name string `json:"name"`
	//Meals interface{} `json:"meals"`
	WeeklyPlan map[string]DailyPlan `json:"weekly_plan"`
	Diet       []string             `json:"diet"`
	Exclusions []string             `json:"exclusions"`
}

type DailyPlan struct {
	Breakfast Meal               `json:"breakfast"`
	Lunch     Meal               `json:"lunch"`
	Dinner    Meal               `json:"dinner"`
	Nutrition map[string]float32 `json:"nutrients"`
}

type Meal struct {
	RecipeID    int    `json:"recipe_id"`
	RecipeTitle string `json:"recipe_title"`
	RecipeImage string `json:"recipe_image"`
	CookTime    int    `json:"ready_in_minutes"`
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

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

func temp(w http.ResponseWriter, req *http.Request) {
	var user User

	f := firego.New("https://souschef-182502.firebaseio.com", nil)
	f.Auth("4maH9UaAtODP5C64FUCpn51Y6kaKSjeSCIHuPZ5y")

	if err := f.Child("users/1").Value(&user); err != nil {
		fmt.Print("ERROR OCCURRED \n", err)
	}

	fmt.Print(json.NewEncoder(w).Encode(user))

	//
	//if err := f.Child("users/1/meals/noday").Set(x); err != nil {
	//	fmt.Print("SOME ERROR OCCURRED\n", err)
	//}
}
