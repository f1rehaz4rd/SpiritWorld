package spiritworldterm

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/f1rehaz4rd/SpiritWorld/pkg/spiritworldtcp"
)

const DELEM = ">="

func queueShell(cmd string) {
	spiritworldtcp.QueueCMD("shell" + DELEM + cmd)
}

func handleInput(text string) {

	cmd := strings.Fields(text)[0]
	tmp := strings.Fields(text)[1:]

	var args string
	for i := 0; i < len(tmp); i++ {
		args += tmp[i] + " "
	}

	switch cmd {
	case "shell":
		queueShell(args)
		break
	default:
		fmt.Println("Invalid command")
	}

}

func TerminalStart() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Spirit World C2:")
	fmt.Println("-----------------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "exit" {
			break
		}

		handleInput(text)
	}

}
