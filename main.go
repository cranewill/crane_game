package main

import (
	"crane_lib/core/net"
	"crane_lib/game"
	"crane_lib/logger"
)

func main() {
	println("Hello World!")
	logger.Init()

	err := net.Start(game.GameServer{net.BaseServer{
		Uuid:   "server_1",
		Status: 0,
	}})
	if err != nil {
		panic("game server start failed")
	}

	signChannel := make(chan int, 1)
	for {
		sign := <-signChannel
		switch sign {
		case -1:
			if err := stop(); err != nil {
				continue
			} else {
				break
			}
		default:
			continue
		}
	}
}

func stop() error {

	return nil
}
