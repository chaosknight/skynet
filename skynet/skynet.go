package skynet

import (
	"sync"
	// "fmt"
	"log"
	"runtime"
	"skynet/types"
	"sync/atomic"
)

type SkyNet struct {
	initialized  bool
	initOptions  types.SkyNetInitOptions
	masterChanel chan types.MasterMsg
	fullChanel   chan types.MasterMsg
	cells        map[string]types.Cell
	callindex    uint64
	msgcount     uint64
	repChanels   sync.Map
}

func (skynet *SkyNet) Init(options types.SkyNetInitOptions) {
	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化初始参数
	if skynet.initialized {
		log.Fatal("请勿重复初始化引擎")
	}
	options.Init()
	skynet.initOptions = options
	skynet.callindex = 0
	skynet.msgcount = 0
	skynet.initialized = true
	skynet.cells = make(map[string]types.Cell)
	skynet.masterChanel = make(chan types.MasterMsg, options.MasterBufferLength)
	skynet.fullChanel = make(chan types.MasterMsg, options.MasterBufferLength)
	for shard := 0; shard < options.MasterSize; shard++ {
		go skynet.masterWorker()
	}

}

func (skynet *SkyNet) Rigist(cell types.Cell, threadsize int) {
	cell.Init(skynet)
	name := cell.GetName()
	_, ok := skynet.cells[name]
	if ok {
		log.Println(name, " cell 已经加载,请勿重复加载")
	} else {
		skynet.cells[name] = cell

		for i := 0; i < threadsize; i++ {
			go cell.Worker()
		}

	}
}

func (skynet *SkyNet) SendMsg(cellname string, cmd string, msgs ...interface{}) {
	skynet.send(false, cellname, cmd, msgs)
}

func (skynet *SkyNet) Call(cellname string, cmd string, msgs ...interface{}) interface{} {
	c := make(chan interface{})
	rid := skynet.send(true, cellname, cmd, msgs)
	if rid == 0 {
		return uint64(0)
	}
	skynet.repChanels.Store(rid, c)
	result := <-c
	return result
}

func (skynet *SkyNet) ReturnResult(cid uint64, result interface{}) {
	if cid > 0 {
		c, _ := skynet.repChanels.Load(cid)
		skynet.repChanels.Delete(cid)
		if cc, ok := c.(chan interface{}); ok {
			cc <- result
		}

	}
	u := atomic.AddUint64(&skynet.msgcount, 1)
	if u == 0 {
		u = atomic.AddUint64(&skynet.msgcount, 1)
	}

}

func (skynet *SkyNet) send(iscall bool, cellname string, cmd string, msgs []interface{}) uint64 {
	if uint(len(skynet.masterChanel)) == skynet.initOptions.MasterBufferLength {
		log.Println(" masterChanel 已经过载")
		return uint64(0)
	}

	rid := skynet.nowindexid()

	if !iscall {
		rid = uint64(0)
	}
	skynet.masterChanel <- types.MasterMsg{Rep: rid, Sid: cellname, Cmd: cmd, Args: msgs}
	return rid
}

func (skynet *SkyNet) nowindexid() uint64 {
	u := atomic.AddUint64(&skynet.callindex, 1)
	if u == 0 {
		u = atomic.AddUint64(&skynet.callindex, 1)
	}
	return u
}

// 阻塞等待直到所有消息完毕
func (skynet *SkyNet) FlushMsg() {
f1:
	for {
		runtime.Gosched()
		if len(skynet.masterChanel) == 0 {
			for _, v := range skynet.cells {
				if len(v.CellChanel()) != 0 {
					continue f1
				}
			}

			if skynet.msgcount == skynet.callindex {
				break f1
			}

		}
	}
}
