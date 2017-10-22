package mas

import (
	"net/http"
	"encoding/json"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"gopkg.in/zabawaba99/firego.v1"
)

var (
	fireURL   = os.Getenv("FIREBASE_URL")
	fireToken = os.Getenv("FIREBASE_AUTH_TOKEN")
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

func writeWeeklyPlanToUser(req *http.Request, wp WeekPlan) error {
	userID := req.URL.Query().Get("user_id")

	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	f.Auth(fireToken)

	err := f.Child("users/" + userID + "/weeklyPlan").Set(wp.Days)
	return err
}
