package main

import (
	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/listeners/tcp"
	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/restapi"
)

func main() {
	go tcp.StartListener()
	restapi.Start()
}
