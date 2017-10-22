package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/karlseguin/ccache"
)

var (
	cache = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

// HANDLERS FOR USER ENDPOINTS

// Will take in a user ID, fetch his/her meal plan for the week from Firebase and return
func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	wp, err := getWeeklyPlanForUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(wp)
}

// Will take in a user ID, fetch his/her shopping list from Firebase and return
func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	user, err := getUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(user)
}

// Will take in a user ID, fetch his/her profile, get a weekly plan from the API, and save to Firebase
func handleCreateWeeklyPlan(w http.ResponseWriter, req *http.Request) {

	if err := createWeeklyPlanForUser(req); err != nil {
		fmt.Fprintln(w, err)
	}

}

// Will take in comma separated list of recipe IDs chosen by user, fetch the ingredients,
// do unit conversions, create shopping list, and save to Firebase
func handleCreateShoppingList(w http.ResponseWriter, req *http.Request) {
	recipeIDs := strings.Split(req.URL.Query().Get("recipe_ids"), ",")

	fmt.Fprint(w, "Generating shopping list for recipes ", recipeIDs)
}

// HANDLERS FOR RECIPE ENDPOINTS

// Will take in a recipeID, get the instructions for it from the API, and return
func handleGetRecipeSteps(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	var steps []Instruction

	recipe_cached := cache.Get("recipe_id:" + recipeID)

	if recipe_cached == nil {
		recipe, err := getRecipeDetails(req)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		cache.Set("recipe_id:"+recipeID, recipe, time.Hour*1000)
		steps = recipe.Instructions
	} else {
		steps = recipe_cached.Value().(Recipe).Instructions
	}

	json.NewEncoder(w).Encode(steps)
}

// Gets ingredients, dietary restrictions, servings, name, instructions, etc.
func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	var recipe Recipe
	recipe_cached := cache.Get("recipe_id:" + recipeID)

	if recipe_cached == nil {
		recipe_to_cache, err := getRecipeDetails(req)

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		cache.Set("recipe_id:"+recipeID, recipe_to_cache, time.Hour*1000)
		recipe = recipe_to_cache
	} else {
		recipe = recipe_cached.Value().(Recipe)
	}

	json.NewEncoder(w).Encode(recipe)

}

// Returns recipe for Leek & Cheese Pie. Static placeholder recipe for testing
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
