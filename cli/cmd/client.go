package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/f1rehaz4rd/SpiritWorld/cli/pkg/agents"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ListAgents", Description: "List all of the agents"},
		{Text: "GetAgent", Description: "<AGENT ID> | Get more information about an agent"},
		{Text: "ListActions", Description: "List all the actions"},
		{Text: "GetAction", Description: "<ACTION ID> | Get more information about an action"},
		{Text: "help", Description: "Show all the commands"},
		{Text: "exit", Description: "Exits the cli"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Welcome to Spirit World Command and Control")
	fmt.Println("Use ctrl+l to clear the terminal window")

	exitFlag := false
	for {
		t := prompt.Input("spiritcli$ ", completer)

		tmp := strings.Split(t, " ")
		cmd := tmp[0]
		switch cmd {
		case "ListAgents":
			agents.ListAgents()
		case "GetAgent":
			agents.GetAgent(tmp[1])
			break
		case "ListActions":
			agents.ListActions()
			break
		case "GetAction":
			agents.GetAction(tmp[1])
			break
		case "exit":
			exitFlag = true
		default:
			fmt.Println("Invalid command")
		}

		if exitFlag {
			break
		}
	}

	// agents.ListAgents()
	// agents.GetAgent("2ee8686e-fcb3-47c6-bae8-7c52df5c4947")
}
