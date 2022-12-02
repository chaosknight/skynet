package skynet

import (
	"log"
	"runtime"
	"sync/atomic"

	"github.com/chaosknight/skynet/types"
)

type SkyNet struct {
	initialized  bool
	initOptions  types.SkyNetInitOptions
	masterChanel chan *types.MasterMsg
	cells        map[string]types.Cell
	callindex    uint64
	msgcount     uint64
	workerSize   int
	isdebug      bool
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
	skynet.isdebug = options.IsDebug
	skynet.workerSize = options.WorkerSize
	skynet.cells = make(map[string]types.Cell)
	skynet.masterChanel = make(chan *types.MasterMsg, options.MasterBufferLength)
	go skynet.masterWorker()

}

func (skynet *SkyNet) Rigist(cell types.Cell, threadsize int) {
	if threadsize == 0 {
		threadsize = skynet.workerSize
	}
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
	skynet.send(nil, cellname, cmd, msgs)
}

func (skynet *SkyNet) Call(cellname string, cmd string, msgs ...interface{}) interface{} {
	c := make(chan interface{})
	ok := skynet.send(c, cellname, cmd, msgs)
	if ok == 0 {
		return uint64(0)
	}
	result := <-c
	return result
}

func (skynet *SkyNet) ReturnResult(msg *types.MasterMsg, result interface{}) {

	if msg.Rep != nil {
		msg.Rep <- result
	}

	u := atomic.AddUint64(&skynet.msgcount, 1)
	if u == 0 {
		u = atomic.AddUint64(&skynet.msgcount, 1)
	}

}

func (skynet *SkyNet) send(cc chan interface{}, cellname string, cmd string, msgs []interface{}) uint64 {
	if uint(len(skynet.masterChanel)) == skynet.initOptions.MasterBufferLength {
		log.Println(" masterChanel 已经过载")
		return uint64(0)
	}
	rid := skynet.nowindexid()
	skynet.masterChanel <- &types.MasterMsg{Rep: cc, Sid: cellname, Cmd: cmd, Args: msgs}
	return rid
}

func (skynet *SkyNet) nowindexid() uint64 {
	u := atomic.AddUint64(&skynet.callindex, 1)
	if u == 0 {
		u = atomic.AddUint64(&skynet.callindex, 1)
	}
	return u
}

func (skynet *SkyNet) Status() []string {
	keys := make([]string, len(skynet.cells))
	j := 0
	for k := range skynet.cells {
		keys[j] = k
		j++
	}
	return keys
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
