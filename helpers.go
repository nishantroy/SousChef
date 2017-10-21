package mas

import (

	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"gopkg.in/zabawaba99/firego.v1"
)

func getUser(req *http.Request) (User, error) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results
	ctx := appengine.NewContext(req)
	client := urlfetch.Client(ctx)
	f := firego.New(fireURL, client)

	var user User

	f.Auth(fireToken)

	err := f.Child("users/" + userID).Value(&user)
	return user, err

}
