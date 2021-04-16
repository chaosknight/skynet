package types

type SkyNetInterface interface {
	SendMsg(cellname string, cmd string, msgs ...interface{})
	Call(cellname string, cmd string, msgs ...interface{}) interface{}
	ReturnResult(msg *MasterMsg, result interface{})
	Status() []string
}
type MasterMsg struct {
	Rep  chan interface{}
	Sid  string
	Cmd  string
	Args []interface{}
}

type CallResp struct {
}

type Cell interface {
	Init(skynet SkyNetInterface)
	Worker() interface{}
	GetName() string
	CellSize() uint
	CellChanel() chan *MasterMsg
	GetSkynet() SkyNetInterface
}
