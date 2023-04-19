package actor

import (
	"github.com/chaosknight/skynet/types"
)

type reducerActor struct {
	name    string
	size    uint
	mchan   chan *types.MasterMsg
	reducer func(types.Actor, *types.MasterMsg)
	skynet  types.SkyNetInterface
}

func NewFromReducer(name string, size uint, reducer func(types.Actor, *types.MasterMsg)) types.Actor {
	ga := &reducerActor{
		name:    name,
		size:    size,
		mchan:   make(chan *types.MasterMsg, size),
		reducer: reducer,
	}
	ga.start()
	return ga
}

func (a *reducerActor) start() {
	go a.receiveLoop()
}

func (a *reducerActor) receiveLoop() {
	for msg := range a.mchan {
		a.reducer(a, msg)
		a.skynet.ReducerCount(msg)
	}
}

func (a *reducerActor) SetMaster(skynet types.SkyNetInterface) {
	a.skynet = skynet
}

func (a *reducerActor) GetSkynet() types.SkyNetInterface {
	return a.skynet
}

func (a *reducerActor) Recive(msg *types.MasterMsg) {
	a.mchan <- msg
}

func (cell *reducerActor) GetName() string {
	return cell.name
}
func (cell *reducerActor) CellSize() uint {
	return cell.size
}
func (cell *reducerActor) CellChanel() chan *types.MasterMsg {
	return cell.mchan
}
