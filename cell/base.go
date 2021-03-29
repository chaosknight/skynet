package cell

import (
	"log"
	"reflect"

	"github.com/chaosknight/skynet/types"
)

type BaseCell struct {
	name   string
	size   uint
	mchan  chan types.MasterMsg
	cmd    map[string]interface{}
	skynet types.SkyNetInterface
}

func (cell *BaseCell) Init(skynet types.SkyNetInterface) {
	cell.mchan = make(chan types.MasterMsg, cell.size)
	cell.cmd = make(map[string]interface{})
	cell.skynet = skynet
	cell.Command("Ping", cell.Ping)
}

func (cell *BaseCell) Command(k string, fun interface{}) {
	cell.cmd[k] = fun
}

func (cell *BaseCell) Ping(msg string) string {
	log.Println(name, " cell recive :", msg)
	return msg
}

func (cell *BaseCell) Worker() interface{} {
	for {
		msg := <-cell.mchan
		log.Println("处理消息:", msg)
		fun, ok := cell.cmd[msg.Cmd]
		if ok {
			result := invoke(fun, msg.Args...)
			log.Println("结果:", result)
			cell.skynet.ReturnResult(msg.Rep, result)
		} else {
			log.Fatal("Command ", msg.Cmd, " is not found ")
			cell.skynet.ReturnResult(msg.Rep, 0)
		}

	}
}

func invoke(fun interface{}, args ...interface{}) interface{} {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	f := reflect.ValueOf(fun)

	v := f.Call(inputs)
	result := []interface{}{}
	for _, vv := range v {
		result = append(result, vv.Interface())

	}
	if len(result) > 0 {
		return result[0]
	}

	return nil
}

func (cell *BaseCell) GetName() string {
	return cell.name
}
func (cell *BaseCell) CellSize() uint {
	return cell.size
}
func (cell *BaseCell) CellChanel() chan types.MasterMsg {
	return cell.mchan
}

func NewCell(name string, size uint) *BaseCell {
	ce := &BaseCell{name: name, size: size}
	return ce
}
