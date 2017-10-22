package mas

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/karlseguin/ccache"
	"bytes"
)

var (
	fireURL    = os.Getenv("FIREBASE_URL")
	fireToken  = os.Getenv("FIREBASE_AUTH_TOKEN")
	spoonToken = os.Getenv("SPOONACULAR_AUTH_TOKEN")
	cache      = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the SousChef API!")

}

// HANDLERS FOR USER ENDPOINTS

// Will take in a user ID, fetch his/her meal plan for the week from Firebase and return
func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	user, err := getUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	json.NewEncoder(w).Encode(user)
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
	user, err := getUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
	}

	diet := strings.Join(user.Diet, ",")
	exclusions := strings.Join(user.Exclusions, ",")

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/mealplans/generate?"
	url += "diet=" + diet + "&exclusions=" + exclusions + "&timeFrame=week"

	//fmt.Fprintln(w, "URL: ", url)

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Fprint(w, "ERROR: ", err)
	}

	request.Header.Set("X-Mashape-Key", spoonToken)

	res, err := client.Do(request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	//fmt.Fprintln(w, s, "\n")

	if err != nil {
		fmt.Print("ERROR: ", err)
	}

	var wp WeekPlan

	defer res.Body.Close()

	json.Unmarshal(buf.Bytes(), &wp)

	//fmt.Fprintln(w, wp)

	json.NewEncoder(w).Encode(wp)

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

	// Make call to API or to Database here, and then write out results

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/"+ recipeID + "/analyzedInstructions"

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

	var steps []Instruction

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&steps)

	json.NewEncoder(w).Encode(steps)
}

// Gets ingredients, dietary restrictions, servings, name, instructions, etc.
func handleGetRecipeDetails(w http.ResponseWriter, req *http.Request) {
	recipeID := req.URL.Query().Get("recipe_id")

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/" + recipeID + "/information"
	//url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/informationBulk?ids=" + recipeID + "&includeNutrition=true"

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

	var recipe Recipe

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&recipe)

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
