package main

import (
	"github.com/f1rehaz4rd/SpiritWorld/pkg/spiritworldtcp"
	"github.com/f1rehaz4rd/SpiritWorld/pkg/spiritworldterm"
)

// const SERVER_ADDR = "127.0.0.1"
const PORT = 2048

func main() {

	go spiritworldtcp.StartServer(PORT)

	spiritworldterm.TerminalStart()
}
