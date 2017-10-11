package mas

import (
	"fmt"
	"net/http"
)



func init() {
	http.HandleFunc("/api/v1/weekly_plan", handleGetWeeklyPlan)
	http.HandleFunc("/", handler)


	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}
