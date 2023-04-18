package types

type SkyNetInterface interface {
	SendMsg(cellname string, cmd string, msgs ...interface{})
	Call(cellname string, cmd string, msgs ...interface{}) interface{}
	ReducerCount(msg *MasterMsg)
	Status() []string
}
type MasterMsg struct {
	Rep  chan interface{}
	Sid  string
	Cmd  string
	Args []interface{}
}

type Actor interface {
	GetName() string
	CellSize() uint
	CellChanel() chan *MasterMsg
	Recive(msg *MasterMsg)
	SetMaster(skynet SkyNetInterface)
}
