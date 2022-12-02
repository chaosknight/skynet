package main

import (
	"github.com/chaosknight/skynet/skynet"
	"github.com/chaosknight/skynet/types"
)

func main() {
	skynet := skynet.SkyNet{}
	skynet.Init(types.SkyNetInitOptions{
		IsDebug: true,
	})
	skynet.Rigist(NewScal(true), 1)

	// skynet.SendMsg(cellScalName, "Ping", "123")
	skynet.FlushMsg()
}
