package main

import (
	"log"

	"github.com/chaosknight/skynet/skynet"
	"github.com/chaosknight/skynet/types"

	"github.com/chaosknight/skynet/cell"
)

func main() {
	skynet := skynet.SkyNet{}
	skynet.Init(types.SkyNetInitOptions{})
	skynet.Rigist(cell.NewAICell(cell.CellAIName, uint(100)), 1)

	// skynet.Rigist(cell.NewMCTreeCell("db", uint(100)), 4)
	// skynet.Rigist(cell.NewCardsCell("cards", uint(100)), 4)

	// first := util.AddRoot(int8(2), [3][15]int8{
	// 	{0, 0, 0, 0, 0, 2, 1, 0, 2, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// })

	// skynet.Call("db", "AddNodes", first)
	// skynet.Call("cards", "AutoPlay", first)
	// set := data.MACDfilter(contrib.LoadLibSVMDataset("C:\\new_zhzq_v6\\T0002\\export\\300059.txt", true))
	// set1 := data.MACDfilter(contrib.LoadLibSVMDataset("C:\\new_zhzq_v6\\T0002\\export\\600196.txt", true))
	// set.AddSet(set1)
	// skynet.Call("ai", "Train", *set)
	// v := util.NewVector(16)
	// v.SetValues([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0})
	// skynet.Call("ai", "ForwardPass", v)
	skynet.SendMsg(cell.CellAIName, "Ping", "569")
	skynet.SendMsg(cell.CellAIName, "Ping", "569")

	skynet.SendMsg(cell.CellAIName, "Ping")
	v := skynet.Call(cell.CellAIName, "Ping")
	log.Println("ddd", v)
	skynet.FlushMsg()
}
