package agents

import (
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

	fmt.Printf("%-16s | %-15s | %-37s | %-21s | %s\n",
		"AGENT NAME",
		"PRIMARY IP",
		"AGENT ID",
		"PUBLIC IP",
		"BEACON LAST SEEN")

	dashHolder := "------------------------------------------------"
	fmt.Printf("%-16s+%-15s+%-37s+%-21s+%s\n",
		dashHolder[:17],
		dashHolder[:17],
		dashHolder[:39],
		dashHolder[:23],
		dashHolder[:38])

	for i := 0; i < len(agents); i++ {
		fmt.Printf("%-16s | %-15s | %-37s | %-21s | %s\n",
			agents[i].AgentObj.AgentName,
			agents[i].AgentObj.PrimaryIP,
			agents[i].AgentObj.UUID,
			agents[i].Publicip,
			agents[i].Beacon.LastBeaconTime)
	}

	fmt.Println()
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

	fmt.Printf("%s - %s\n\n", agent.AgentObj.AgentName, agent.AgentObj.UUID)
	fmt.Printf("\tAgentVersion: %s\n", agent.AgentObj.AgentVersion)
	fmt.Printf("\tPrimary IP: %s\n", agent.AgentObj.PrimaryIP)
	fmt.Printf("\tHostname: %s\n", agent.AgentObj.Hostname)
	fmt.Printf("\tMAC: %s\n", agent.AgentObj.MAC)
	fmt.Printf("\tOS: %s\n", agent.AgentObj.AgentOS)
	fmt.Printf("\tOther IPs: %s\n", agent.AgentObj.OtherIPs)
	fmt.Printf("\tPublic IP: %s\n\n", agent.Publicip)
	fmt.Printf("\tRegister Time: %s\n", agent.Beacon.RegisterTime)
	fmt.Printf("\tLast Beacon Time: %s\n", agent.Beacon.LastBeaconTime)
	fmt.Printf("\tAction Queue: \n\t%s\n\n", agent.Beacon.ActionQueue)
	fmt.Printf("\tActions: \n\t%s\n\n", agent.Beacon.Actions)
}

func ListActions() {
	response, _ := http.Get("http://127.0.0.1:8080/api/actions")
	data, _ := ioutil.ReadAll(response.Body)

	var actions []models.ActionModel

	json.Unmarshal(data, &actions)

	fmt.Printf("%-36s | %-36s | %-13s | %-21s\n",
		"ACTION UUID",
		"AGENT UUID",
		"ACTION TYPE",
		"ACTION CMD")

	dashHolder := "------------------------------------------------"
	fmt.Printf("%-37s+%-37s+%-13s+%-21s\n",
		dashHolder[:37],
		dashHolder[:38],
		dashHolder[:15],
		dashHolder[:23])

	for i := 0; i < len(actions); i++ {
		fmt.Printf("%-36s | %-36s | %-13s | %s\n",
			actions[i].UUID,
			actions[i].AgentUUID,
			actions[i].ActionType,
			actions[i].ActionCmd)
	}

	fmt.Println()
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

	fmt.Printf("%s\n\n", action.UUID)
	fmt.Printf("\tAgent UUID: %s\n", action.AgentUUID)
	fmt.Printf("\tAction Type: %s\n", action.ActionType)
	fmt.Printf("\tAction Command: %s\n", action.ActionCmd)
	fmt.Printf("\tAction Output: \n\t%s\n", action.ActionOutput)
}
