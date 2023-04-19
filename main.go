package main

import (
	"log"

	"github.com/chaosknight/skynet/actor"
	"github.com/chaosknight/skynet/skynet"
	"github.com/chaosknight/skynet/types"
)

func main() {
	skynet := skynet.SkyNet{}
	skynet.Init(types.SkyNetInitOptions{
		IsDebug: true,
	})

	testactor := actor.NewFromReducer("Custom", 10, func(a types.Actor, msg *types.MasterMsg) {
		log.Println(" cell recive :", msg.Cmd, msg.Args)
		if msg.Rep != nil {
			msg.Rep <- "999"
		}

	})
	skynet.Rigist(testactor, 1)

	results := skynet.Call("Custom", "Custom", "123")
	log.Println(" cell recive :", results)
	skynet.FlushMsg()
}
