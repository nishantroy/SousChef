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
	if err := f.Child("users/" + userID).Set(user); err != nil {
		return err
	}
	return nil
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
	if err := f.Child("users/" + userID).Set(user); err != nil {
		return err
	}
	return nil
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

	if err = writeWeeklyPlanToUser(req, wp); err != nil {
		return err
	}

	return nil
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
		this_meal := MealTemp{ID: id, Name: name}

		var day_update Day

		day_update = wp.Days[day]

		switch mealnumber {
		case 1:
			day_update.Breakfast = this_meal
		case 2:
			day_update.Lunch = this_meal
		default:
			day_update.Dinner = this_meal
		}

		wp.Days[day] = day_update
	}

	return nil
}
