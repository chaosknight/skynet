package main

import (
	"github.com/chaosknight/skynet/skynet"
	"github.com/chaosknight/skynet/types"
)

func main() {
	skynet := skynet.SkyNet{}
	skynet.Init(types.SkyNetInitOptions{})
	// skynet.Rigist(NewScal(), 1)

	// skynet.SendMsg(cellScalName, "Ping", "123")
	skynet.FlushMsg()
}
