package cell

import (
	"log"
	"reflect"

	"github.com/chaosknight/skynet/types"
)

type Cell struct {
	isdebug bool
	name    string
	size    uint
	mchan   chan *types.MasterMsg
	cmd     map[string]reflect.Value
	skynet  types.SkyNetInterface
}

func (cell *Cell) Init(skynet types.SkyNetInterface) {
	cell.mchan = make(chan *types.MasterMsg, cell.size)
	cell.cmd = make(map[string]reflect.Value)
	cell.skynet = skynet
	cell.command(cell)

}

func (cell *Cell) command(c types.Cell) {
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
		if cell.isdebug {
			log.Println(typ.Method(i).Name)
		}

	}
}

func (cell *Cell) GetSkynet() types.SkyNetInterface {
	return cell.skynet
}

func (cell *Cell) Ping(msg string) string {
	log.Println(cell.name, " cell recive :", msg)
	return msg
}

func (cell *Cell) Worker() interface{} {
	for {
		msg := <-cell.mchan
		fun, ok := cell.cmd[msg.Cmd]
		if ok {
			result := invoke(fun, msg.Args...)
			cell.skynet.ReturnResult(msg, result)
		} else {
			log.Fatal("Command ", msg.Cmd, " is not found ")
			cell.skynet.ReturnResult(msg, 0)
		}
	}
}

func invoke(fun reflect.Value, args ...interface{}) (res interface{}) {
	defer func() {
		if x := recover(); x != nil {
			log.Println("call work error", x)
			res = nil
		}
	}()
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

func (cell *Cell) GetName() string {
	return cell.name
}
func (cell *Cell) CellSize() uint {
	return cell.size
}
func (cell *Cell) CellChanel() chan *types.MasterMsg {
	return cell.mchan
}

func GO(name string, size uint, isdebug bool) *Cell {
	ce := &Cell{name: name, size: size, isdebug: isdebug}
	return ce
}
