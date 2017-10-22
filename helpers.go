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

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err = f.Child("users/" + userID + "/weekly_plan/" + day + "/" + meal).Set(MealTemp{ID: recipeID, Name: recipeName})
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

	return recipe, nil
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

		id := int(value["id"].(float64))
		name := value["title"].(string)
		thisMeal := MealTemp{ID: id, Name: name}

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
