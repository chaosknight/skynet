package skynet

import (
	"log"

	"github.com/chaosknight/skynet/types"
)

func (skynet *SkyNet) masterWorker() {

	for {
		msg := <-skynet.masterChanel
		chanel := skynet.getCellChanel(msg.Sid)
		if chanel != nil {
			chanel <- msg
		}

	}
}

func (skynet *SkyNet) getCellChanel(name string) chan *types.MasterMsg {
	v, ok := skynet.cells[name]
	if ok {
		return v.CellChanel()
	} else {
		log.Fatal("cell ", name, " is not found ")
		return nil
	}

}
