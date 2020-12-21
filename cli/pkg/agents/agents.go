package agents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/cli/pkg/models"
)

func ListAgents() {
	response, _ := http.Get("http://127.0.0.1:8080/api/agents")
	data, _ := ioutil.ReadAll(response.Body)

	var agents []models.AgentModel
	json.Unmarshal(data, &agents)
	models.PrintAgents(agents)
}

func GetAgent(id string) {
	response, _ := http.Get("http://127.0.0.1:8080/api/agent/" + id)
	data, _ := ioutil.ReadAll(response.Body)

	var agent models.AgentModel
	json.Unmarshal(data, &agent)
	if agent.AgentObj == nil {
		fmt.Println("Agent ID is invalid")
		return
	}

	models.PrintAgent(agent)
}

func ListActions() {
	response, _ := http.Get("http://127.0.0.1:8080/api/actions")
	data, _ := ioutil.ReadAll(response.Body)

	var actions []models.ActionModel
	json.Unmarshal(data, &actions)
	models.PrintActions(actions)
}

func GetAction(id string) {
	response, _ := http.Get("http://127.0.0.1:8080/api/action/" + id)
	data, _ := ioutil.ReadAll(response.Body)

	var action models.ActionModel
	json.Unmarshal(data, &action)
	if action.ActionType == "" {
		fmt.Println("Action ID is invalid")
		return
	}

	models.PrintAction(action)
}

func CreateAction(id, actionType, cmd string) {
	urlStr := "http://localhost:8080/api/createaction/" + id

	action := models.ActionModel{
		ActionType: actionType,
		ActionCmd:  cmd,
		AgentUUID:  id,
	}

	jsonStr, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var responseAction models.ActionModel
	err = json.Unmarshal(body, &responseAction)
	if err != nil {
		fmt.Println(err)
	}

	models.PrintAction(responseAction)
}

func CreateGroupAction(name, actionType, cmd string) {
	urlStr := "http://localhost:8080/api/creategroupaction/" + name

	action := models.ActionModel{
		ActionType: actionType,
		ActionCmd:  cmd,
	}

	jsonStr, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Didn't error")
}

func ListGroups() {
	response, _ := http.Get("http://127.0.0.1:8080/api/groups")
	data, _ := ioutil.ReadAll(response.Body)

	var groups []models.GroupModel
	json.Unmarshal(data, &groups)
	models.PrintGroups(groups)
}

func GetGroup(name string) {
	response, _ := http.Get("http://127.0.0.1:8080/api/group/" + name)
	data, _ := ioutil.ReadAll(response.Body)

	var group models.GroupModel
	json.Unmarshal(data, &group)
	if group.GroupName == "" {
		fmt.Println("Group name is invalid")
		return
	}

	models.PrintGroup(group)
}

func CreateGroup(name string) {
	urlStr := "http://localhost:8080/api/creategroup/" + name

	group := models.GroupModel{
		GroupName: name,
	}

	jsonStr, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var responseGroup models.GroupModel
	err = json.Unmarshal(body, &responseGroup)
	if err != nil {
		fmt.Println(err)
	}

	models.PrintGroup(responseGroup)
}

func AddToGroup(name, id string) {
	urlStr := "http://localhost:8080/api/addtogroup/" + name + "/" + id

	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func RemoveFromGroup(name, id string) {
	urlStr := "http://localhost:8080/api/removefromgroup/" + name + "/" + id

	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
