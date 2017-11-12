package mas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Takes in a user ID, fetches meal plan for the week from Firebase and returns
func handleGetWeeklyPlan(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wp, err := getWeeklyPlanForUser(req)

	if err != nil {
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wp.Days)
}

// Takes in a user ID, fetches profile, gets a weekly plan from the API, and saves to Firebase
func handleCreateWeeklyPlan(w http.ResponseWriter, req *http.Request) {

	if err := createWeeklyPlanForUser(req); err != nil {
		fmt.Fprintln(w, err)
	}

}

func handleUpdateMeal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateMeal(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "ERROR IN HELPER METHOD :", err)
		return
	}
}

func handleGetRecipeForAlexa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rec, err := GetRecipeForAlexa(w, req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("error occurred")
		return
	}

	json.NewEncoder(w).Encode(rec)
}
