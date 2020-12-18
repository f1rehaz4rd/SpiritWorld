package main

import (
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/RestAPI"
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/listeners/tcp"
)

func main() {
	go tcp.StartListener()
	RestAPI.Start()
}
