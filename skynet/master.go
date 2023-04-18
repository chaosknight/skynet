package skynet

import (
	"log"

	"github.com/chaosknight/skynet/types"
)

func (skynet *SkyNet) masterWorker() {

	for {
		msg := <-skynet.masterChanel
		actor := skynet.getActor(msg.Sid)
		if actor != nil {
			actor.Recive(msg)
		}
	}
}

func (skynet *SkyNet) getActor(name string) types.Actor {
	v, ok := skynet.cells[name]
	if ok {
		return v
	} else {
		log.Fatal("actor ", name, " is not found ")
		return nil
	}

}
