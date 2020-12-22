package RestAPI

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// Init Router
	r := mux.NewRouter()

	db.Open()
	defer db.Close()

	r.HandleFunc("/api/agents", getAgents).Methods("GET")
	r.HandleFunc("/api/agent/{id}", getAgent).Methods("GET")
	r.HandleFunc("/api/actions", getActions).Methods("GET")
	r.HandleFunc("/api/action/{id}", getAction).Methods("GET")
	r.HandleFunc("/api/groups", getGroups).Methods("GET")
	r.HandleFunc("/api/group/{name}", getGroup).Methods("GET")
	r.HandleFunc("/api/createaction/{id}", createAction).Methods("POST")
	r.HandleFunc("/api/creategroupaction/{name}", createGroupAction).Methods("POST")
	r.HandleFunc("/api/creategroup/{name}", createGroup).Methods("POST")
	r.HandleFunc("/api/addtogroup/{name}/{id}", addToGroup).Methods("POST")
	r.HandleFunc("/api/removefromgroup/{name}/{id}", removeFromGroup).Methods("POST")

	fmt.Println("Starting RestAPI on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
