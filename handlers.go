package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

// HANDLERS FOR USER ENDPOINTS

// Takes in a user ID, fetches meal plan for the week from Firebase and returns
func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wp, err := getWeeklyPlanForUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wp.Days)
}

// Takes in a user ID, fetches profile, gets a weekly plan from the API, and saves to Firebase
func handleCreateWeeklyPlan(w http.ResponseWriter, req *http.Request) {

	if err := createWeeklyPlanForUser(req); err != nil {
		fmt.Fprintln(w, err)
	}

}

// Will take in a user ID, fetch his/her shopping list from Firebase and return
func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := getUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Will take in comma separated list of recipe IDs chosen by user, fetch the ingredients,
// do unit conversions, create shopping list, and save to Firebase
func handleCreateShoppingList(w http.ResponseWriter, req *http.Request) {
	recipeIDs := strings.Split(req.URL.Query().Get("recipe_ids"), ",")
	fmt.Println(len(recipeIDs))
	for i := 0; i < len(recipeIDs); i++ {
		id := recipeIDs[i]
		r, err := getRecipeDetails(req, id)
		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			fmt.Fprintln(w, r.Ingredients)
		}
	}

	fmt.Fprint(w, "Generating shopping list for recipes ", recipeIDs)
}

func handleCreateProfile(w http.ResponseWriter, req *http.Request) {
	err := createUserProfile(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func handleUpdateProfile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateUserProfile(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

// HANDLERS FOR RECIPE ENDPOINTS
func handleGetRecipeChanges(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recipes, err := getRecipeChanges(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "ERROR IN HELPER METHOD :", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)

}

func handleUpdateMeal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateMeal(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "ERROR IN HELPER METHOD :", err)
		return
	}
}

// Takes in a recipeID, gets the instructions for it from the API/cache, and returns
func handleGetRecipeSteps(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	recipeID := req.URL.Query().Get("recipe_id")

	recipe, err := getRecipeDetails(req, recipeID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe.Instructions)

	//var steps []Instruction
	//
	//recipeCached := cache.Get("recipe_id:" + recipeID)
	//
	//if recipeCached == nil {
	//	recipe, err := getRecipeDetails(req)
	//
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		fmt.Fprintln(w, err)
	//		return
	//	}
	//
	//	cache.Set("recipe_id:"+recipeID, recipe, time.Hour*1000)
	//	steps = recipe.Instructions
	//} else {
	//	steps = recipeCached.Value().(Recipe).Instructions
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(steps)
}

// Takes in a recipeID, gets the details for it from the API/cache, and returns
func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")
	w.Header().Set("Content-Type", "application/json")

	recipe, err := getRecipeDetails(req, recipeID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipe)

	//recipeID := req.URL.Query().Get("recipe_id")
	//
	//var recipe Recipe
	//recipeCached := cache.Get("recipe_id:" + recipeID)
	//
	//if recipeCached == nil {
	//	recipeCached, err := getRecipeDetails(req)
	//
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		fmt.Fprintln(w, err)
	//		return
	//	}
	//
	//	cache.Set("recipe_id:"+recipeID, recipeCached, time.Hour*1000)
	//	recipe = recipeCached
	//} else {
	//	recipe = recipeCached.Value().(Recipe)
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(recipe)

}

/* Returns recipe for Leek & Cheese Pie. Static placeholder recipe for testing
func handleStaticRecipeDetails(w http.ResponseWriter, _ *http.Request) {

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
*/
