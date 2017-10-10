package main

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
)



func main() {
	appengine.Main()
	http.HandleFunc("/api/v1/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/", handler)


	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}