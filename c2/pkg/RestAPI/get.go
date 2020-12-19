package RestAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := db.AllAgents()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(agents)
	}
}

func getAgent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agent, err := db.GetAgentByID(params["id"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(agent)
	}
}

func getActions(w http.ResponseWriter, r *http.Request) {
	actions, err := db.AllActions()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action, err := db.GetActionByID(params["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := db.AllGroups()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(groups)
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	group, err := db.GetGroupByID(params["name"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}
