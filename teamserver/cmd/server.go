package main

import (
	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/RestAPI"
	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/listeners/tcp"
)

func main() {
	go tcp.StartListener()
	RestAPI.Start()
}
