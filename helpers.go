package mas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"errors"
	"github.com/karlseguin/ccache"
	"gopkg.in/zabawaba99/firego.v1"
)

var (
	fireURL    = os.Getenv("FIREBASE_URL")
	fireToken  = os.Getenv("FIREBASE_AUTH_TOKEN")
	spoonToken = os.Getenv("SPOONACULAR_AUTH_TOKEN")
	cache      = ccache.New(ccache.Configure().MaxSize(1000).ItemsToPrune(100))
	replacer   = strings.NewReplacer(".", ",", "$", ",", "[", ",", "]", ",", "#", ",", "/", ",")
)

func getUser(req *http.Request) (User, error) {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	var user User

	f.Auth(fireToken)

	err := f.Child("users/" + userID).Value(&user)
	return user, err

}

func createUserProfile(req *http.Request) error {
	userID := req.URL.Query().Get("user_id")
	userName := req.URL.Query().Get("name")
	diet := req.URL.Query().Get("diet")
	exclusions := strings.Split(req.URL.Query().Get("exclusions"), ",")

	user := User{Name: userName, Diet: diet, Exclusions: exclusions}

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err := f.Child("users/" + userID).Set(user)
	return err

}

func updateUserProfile(req *http.Request) error {
	user, err := getUser(req)

	if err != nil {
		return err
	}

	userID := req.URL.Query().Get("user_id")
	diet := req.URL.Query().Get("diet")
	exclusions := strings.Split(req.URL.Query().Get("exclusions"), ",")

	user.Diet = diet
	user.Exclusions = exclusions

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)
	err = f.Child("users/" + userID).Set(user)
	return err
}

func getWeeklyPlanForUser(req *http.Request) (WeekPlan, error) {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	wp := WeekPlan{}

	f.Auth(fireToken)

	err := f.Child("users/" + userID + "/weekly_plan").Value(&wp.Days)
	return wp, err
}

func createWeeklyPlanForUser(req *http.Request) error {
	user, err := getUser(req)

	if err != nil {
		return err
	}

	fmt.Println(user)

	diet := user.Diet
	exclusions := strings.Join(user.Exclusions, ",")

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/mealplans/generate?"
	url += "diet=" + diet + "&exclusions=" + exclusions + "&timeFrame=week"

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	request.Header.Set("X-Mashape-Key", spoonToken)

	res, err := client.Do(request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	if err != nil {
		fmt.Print("ERROR: ", err)
	}

	var wp WeekPlan

	defer res.Body.Close()

	json.Unmarshal(buf.Bytes(), &wp)

	// Get cook time and images for all meals
	for index, day := range wp.Days {
		var (
			breakfast Recipe
			lunch     Recipe
			dinner    Recipe
		)

		// Get cook time and image for breakfast
		breakfastID := strconv.Itoa(day.Breakfast.ID)
		q := req.URL.Query()
		q.Set("recipe_id", breakfastID)
		req.URL.RawQuery = q.Encode()
		breakfast, err = getRecipeDetails(req, breakfastID)
		day.Breakfast.CookTime = breakfast.CookTime
		day.Breakfast.Image = breakfast.Image

		// Get cook time and image for lunch
		lunchID := strconv.Itoa(day.Lunch.ID)
		q = req.URL.Query()
		q.Set("recipe_id", lunchID)
		req.URL.RawQuery = q.Encode()
		lunch, err = getRecipeDetails(req, lunchID)
		day.Lunch.CookTime = lunch.CookTime
		day.Lunch.Image = lunch.Image

		// Get cook time and image for dinner
		dinnerID := strconv.Itoa(day.Dinner.ID)
		q = req.URL.Query()
		q.Set("recipe_id", dinnerID)
		req.URL.RawQuery = q.Encode()
		dinner, err = getRecipeDetails(req, dinnerID)
		day.Dinner.CookTime = dinner.CookTime
		day.Dinner.Image = dinner.Image

		wp.Days[index] = day

	}

	err = writeWeeklyPlanToUser(req, wp)
	return err
}

func writeWeeklyPlanToUser(req *http.Request, wp WeekPlan) error {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err := f.Child("users/" + userID + "/weekly_plan").Set(wp.Days)
	if err != nil {
		return err
	}

	err = f.Child("users/" + userID + "/weekly_plan_date").Set(time.Now().Format("02-01-2006"))
	return err
}

func updateMeal(req *http.Request) error {
	userID := req.URL.Query().Get("user_id")
	day := req.URL.Query().Get("day")
	meal := req.URL.Query().Get("meal")

	recipeID, err := strconv.Atoi(req.URL.Query().Get("recipe_id"))

	if err != nil {
		return err
	}

	recipe, err := getRecipeDetails(req, strconv.Itoa(recipeID))

	if err != nil {
		return err
	}

	recipeName := recipe.Title
	cookTime := recipe.CookTime
	image := recipe.Image

	newMeal := Meal{ID: recipeID, Name: recipeName, CookTime: cookTime, Image: image}

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err = f.Child("users/" + userID + "/weekly_plan/" + day + "/" + meal).Set(newMeal)
	return err

}

func getRecipeChanges(req *http.Request) (RecipeChanges, error) {
	var r RecipeChanges
	user, err := getUser(req)

	if err != nil {
		return r, err
	}

	offset := req.URL.Query().Get("offset")
	mealType := req.URL.Query().Get("meal_type")

	diet := user.Diet
	exclusions := strings.Join(user.Exclusions, ",")

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/search?"
	url += "diet=" + diet + "&excludeIngredients=" + exclusions + "&number=10" + "&offset=" + offset +
		"&type=" + mealType

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return r, err
	}

	request.Header.Set("X-Mashape-Key", spoonToken)

	res, err := client.Do(request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	if err != nil {
		fmt.Print("ERROR: ", err)
	}

	defer res.Body.Close()

	json.Unmarshal(buf.Bytes(), &r)

	return r, err
}

func getShoppingListForUser(req *http.Request) (map[string]map[string]map[string]float32, error) {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	var shopList map[string]map[string]map[string]float32

	f.Auth(fireToken)

	err := f.Child("users/" + userID + "/shopping_list").Value(&shopList)
	return shopList, err
}

func createShoppingListForUser(req *http.Request) error {
	recipeIDs := strings.Split(req.URL.Query().Get("recipe_ids"), ",")

	var shopList = make(map[string]map[string]map[string]float32)
	for _, id := range recipeIDs {
		r, err := getRecipeDetails(req, id)
		if err != nil {
			return err
		}

		ingredients := r.Ingredients
		for _, ingredient := range ingredients {
			category := replacer.Replace(strings.Split(ingredient.Category, ";")[0])
			unit := replacer.Replace(ingredient.Unit)
			if unit == "" {
				unit = "empty unit"
			}
			name := ingredient.Name
			if name == "water" {
				continue
			}

			_, catExists := shopList[category]
			if !catExists {
				itemMap := make(map[string]map[string]float32)
				shopList[category] = itemMap

			}
			itemMap := shopList[category]

			_, itemExists := itemMap[name]

			if !itemExists {
				unitMap := make(map[string]float32)
				itemMap[name] = unitMap
				shopList[category] = itemMap
			}

			_, unitExists := itemMap[name][unit]

			if !unitExists {
				itemMap[name][unit] = float32(0)
			}
			itemMap[name][unit] += ingredient.Amount

		}

	}

	return writeShoppingListToUser(req, shopList)
}

func writeShoppingListToUser(req *http.Request, shopList map[string]map[string]map[string]float32) error {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err := f.Child("users/" + userID + "/shopping_list").Set(&shopList)
	if err != nil {
		return errors.New("FIREBASE ERROR: " + err.Error())
	}
	return err
}

func getRecipeDetails(req *http.Request, recipeID string) (Recipe, error) {
	recipeCached := cache.Get("recipe_id:" + recipeID)

	if recipeCached == nil {
		var recipe Recipe

		url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/" + recipeID + "/information"

		ctx := appengine.NewContext(req)
		client := urlfetch.Client(ctx)

		request, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return recipe, err
		}

		request.Header.Set("X-Mashape-Key", spoonToken)

		res, err := client.Do(request)

		if err != nil {
			return recipe, err
		}

		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&recipe)

		if recipe.Title == "" {
			return recipe, errors.New("something went wrong, the recipe name is empty")
		}
		cache.Set("recipe_id:"+recipeID, recipe, time.Hour*1000)
		return recipe, nil
	}

	return recipeCached.Value().(Recipe), nil

}

// UnmarshalJSON is overwritten for the WeekPlan struct to handle the nested JSON returned from the API gracefully
func (wp *WeekPlan) UnmarshalJSON(b []byte) error {
	wp.Days = make([]Day, 7)
	var f map[string]*json.RawMessage
	json.Unmarshal(b, &f)

	var v []map[string]interface{}
	json.Unmarshal(*f["items"], &v)

	for _, item := range v {

		day := int(item["day"].(float64)) - 1
		mealnumber := int(item["slot"].(float64))

		var value map[string]interface{}
		json.Unmarshal([]byte(item["value"].(string)), &value)

		fmt.Println(value)

		id := int(value["id"].(float64))
		name := value["title"].(string)

		thisMeal := Meal{ID: id, Name: name}

		var dateUpdate Day

		dateUpdate = wp.Days[day]

		switch mealnumber {
		case 1:
			dateUpdate.Breakfast = thisMeal
		case 2:
			dateUpdate.Lunch = thisMeal
		default:
			dateUpdate.Dinner = thisMeal
		}

		wp.Days[day] = dateUpdate
	}

	return nil
}
