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
	cmd    map[string]reflect.Value
	skynet types.SkyNetInterface
}

func (cell *BaseCell) Init(skynet types.SkyNetInterface, ccell types.Cell) {
	cell.mchan = make(chan types.MasterMsg, cell.size)
	cell.cmd = make(map[string]reflect.Value)
	cell.skynet = skynet
	cell.command(ccell)

}

func (cell *BaseCell) command(c types.Cell) {
	value := reflect.ValueOf(c)
	typ := value.Type()
	for i := 0; i < value.NumMethod(); i++ {
		switch typ.Method(i).Name {
		case "Init":
		case "Worker":
			break
		default:
			cell.cmd[typ.Method(i).Name] = value.Method(i)
		}

	}
}

func (cell *BaseCell) GetSkynet() types.SkyNetInterface {
	return cell.skynet
}

func (cell *BaseCell) Ping(msg string) string {
	log.Println(cell.name, " cell recive :", msg)
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

func invoke(fun reflect.Value, args ...interface{}) interface{} {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	f := fun

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
