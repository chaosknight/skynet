package types

type SkyNetInterface interface {
	SendMsg(cellname string, cmd string, msgs ...interface{})
	Call(cellname string, cmd string, msgs ...interface{}) interface{}
	ReturnResult(cid uint64, result interface{})
}
type MasterMsg struct {
	Rep  uint64
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
	CellChanel() chan MasterMsg
	GetSkynet() SkyNetInterface
}
