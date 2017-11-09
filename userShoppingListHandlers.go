package mas

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// Will take in a user ID, fetch his/her shopping list from Firebase and return
func handleGetShoppingList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shopList, err := getShoppingListForUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shopList)
}

// Will take in comma separated list of recipe IDs chosen by user, fetch the ingredients,
// do unit conversions, create shopping list, and save to Firebase
func handleCreateShoppingList(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := createShoppingListForUser(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintln(w, "SUCCESS!")
}

func handleCheckGroceryItem(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateGroceryItemDoneForUser(req, true)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleUncheckGroceryItem(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := updateGroceryItemDoneForUser(req, false)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "SOME ERROR OCCURRED", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
