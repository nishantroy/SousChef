package mas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"gopkg.in/zabawaba99/firego.v1"
	"strconv"
)

var (
	fireURL    = os.Getenv("FIREBASE_URL")
	fireToken  = os.Getenv("FIREBASE_AUTH_TOKEN")
	spoonToken = os.Getenv("SPOONACULAR_AUTH_TOKEN")
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
	diet := strings.Split(req.URL.Query().Get("diet"), ",")
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
	diet := strings.Split(req.URL.Query().Get("diet"), ",")
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

func updateMeal(req *http.Request) error {
	userID := req.URL.Query().Get("user_id")
	day := req.URL.Query().Get("day")
	meal := req.URL.Query().Get("meal")

	recipeID, err := strconv.Atoi(req.URL.Query().Get("recipe_id"))

	if err != nil {
		return err
	}

	recipeName := req.URL.Query().Get("recipe_name")
	cookTime, err := strconv.Atoi(req.URL.Query().Get("cook_time"))
	if err != nil {
		return err
	}
	image := req.URL.Query().Get("image")

	new_meal := Meal{ID: recipeID, Name: recipeName, CookTime: cookTime, Image: image}

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err = f.Child("users/" + userID + "/weekly_plan/" + day + "/" + meal).Set(new_meal)
	return err

}

func getRecipeChanges(req *http.Request) (RecipeChanges, error) {
	var r RecipeChanges
	user, err := getUser(req)

	if err != nil {
		return r, err
	}

	offset := req.URL.Query().Get("offset")
	meal_type := req.URL.Query().Get("meal_type")

	diet := strings.Join(user.Diet, ",")
	exclusions := strings.Join(user.Exclusions, ",")

	url := "https://spoonacular-recipe-food-nutrition-v1.p.mashape.com/recipes/search?"
	url += "diet=" + diet + "&excludeIngredients=" + exclusions + "&number=10" + "&offset=" + offset +
		"&type=" + meal_type

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

func createWeeklyPlanForUser(req *http.Request) error {
	user, err := getUser(req)

	if err != nil {
		return err
	}

	fmt.Println(user)

	diet := strings.Join(user.Diet, ",")
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
		breakfast_id := strconv.Itoa(day.Breakfast.ID)
		q := req.URL.Query()
		q.Set("recipe_id", breakfast_id)
		req.URL.RawQuery = q.Encode()
		breakfast, err = getRecipeDetails(req)
		day.Breakfast.CookTime = breakfast.CookTime
		day.Breakfast.Image = breakfast.Image

		// Get cook time and image for lunch
		lunch_id := strconv.Itoa(day.Lunch.ID)
		q = req.URL.Query()
		q.Set("recipe_id", lunch_id)
		req.URL.RawQuery = q.Encode()
		lunch, err = getRecipeDetails(req)
		day.Lunch.CookTime = lunch.CookTime
		day.Lunch.Image = lunch.Image

		// Get cook time and image for dinner
		dinner_id := strconv.Itoa(day.Dinner.ID)
		q = req.URL.Query()
		q.Set("recipe_id", dinner_id)
		req.URL.RawQuery = q.Encode()
		dinner, err = getRecipeDetails(req)
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

func getRecipeDetails(req *http.Request) (Recipe, error) {
	recipeID := req.URL.Query().Get("recipe_id")

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

		cache.Set("recipe_id:"+recipeID, recipe, time.Hour*1000)
		return recipe, nil
	} else {
		return recipeCached.Value().(Recipe), nil
	}
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
