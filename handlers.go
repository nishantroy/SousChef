package mas

import (
	"fmt"
	"gopkg.in/zabawaba99/firego.v1"
	"net/http"

	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"os"
)

var (
	fireURL    = os.Getenv("FIREBASE_URL")
	fireToken  = os.Getenv("FIREBASE_AUTH_TOKEN")
	spoonToken = os.Getenv("SPOONACULAR_AUTH_TOKEN")
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	var user User

	f.Auth(fireToken)

	if err := f.Child("users/" + userID).Value(&user); err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(user)
}

func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	var user User

	f.Auth(fireToken)

	if err := f.Child("users/" + userID).Value(&user); err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(user)
}

func handleGetRecipeSteps(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello! The steps for recipe %s are: _____", recipeID)
}

// Gets ingredients, dietary restrictions, servings, name, etc. No steps
func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	//url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/" + recipeID + "/information"
	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/informationBulk?ids=" + recipeID+"&includeNutrition=true"


	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Fprint(w, "ERROR: ", err)
	}

	request.Header.Set("X-Mashape-Key", spoonToken)

	res, err := client.Do(request)

	if err != nil {
		fmt.Print("ERROR: ", err)
	}

	var recipes []Recipe

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&recipes)

	// Make call to API or to Database here, and then write out results

	json.NewEncoder(w).Encode(recipes)

}

func handleStaticRecipeDetails(w http.ResponseWriter, req *http.Request) {

	var recipe Recipe

	recipe.Title = "Leek & Cheese Pie"
	recipe.CookTime = 75
	recipe.ID = 116679
	recipe.Image = "https://spoonacular.com/recipeImages/116679-556x370.jpg"
	recipe.Cheap = false
	recipe.Vegetarian = true
	recipe.Vegan = false
	recipe.Ketogenic = false
	recipe.Servings = 4

	recipe.Ingredients = []Ingredient{
		{ID: 18371, Category: "Baking", Name: "baking powder", Amount: 2, Unit: "tsp",
			FullDescriptor: "2 teaspoons baking powder"},
		{ID: 1001, Category: "Milk, Eggs, Other Dairy", Name: "butter", Amount: 100, Unit: "g",
			FullDescriptor: "100 g butter or 100 g margarine"},
		{ID: 2031, Category: "Spices and Seasonings", Name: "cayenne pepper", Amount: 4,
			Unit: "servings", FullDescriptor: "cayenne pepper"},
	}

	json.NewEncoder(w).Encode(recipe)

}
