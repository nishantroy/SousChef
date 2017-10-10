package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the Sous Chef API!")
}

func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("user_id")

	// Make call to API or to Database here, and then write out results

	fmt.Fprintf(w, "Hello User %s! Your weekly plan is: _____", userID)
}
