package RestAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/database"
	"github.com/gorilla/mux"
)

var db database.DatabaseModel

func createAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	var action database.ActionModel
	err := decoder.Decode(&action)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if !db.InsertAction(params["id"], &action) {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func createGroupAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)
	var action database.ActionModel
	err := decoder.Decode(&action)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	check := db.InsertGroupAction(params["name"], action)
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	group := database.GroupModel{
		GroupName: params["name"],
	}

	w.Header().Set("Content-Type", "application/json")
	if !db.CreateGroup(group) {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func addToGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	check := db.InsertIntoGroup(params["name"], params["id"])
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}

func removeFromGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	check := db.RemoveFromGroup(params["name"], params["id"])
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}
