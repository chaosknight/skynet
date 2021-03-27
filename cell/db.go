package cell

import (
	"log"

	"skynet/types"
)

const DBCellName = "kvdb"

type dbcell struct {
	BaseCell
	initialized bool
}

func (db *dbcell) Init(skynet types.SkyNetInterface) {
	if db.initialized == true {
		log.Fatal("simpleai不能初始化两次")
	}
	db.BaseCell.Init(skynet)
	db.initialized = true
	db.Command("Ping", db.Ping)
}

func (db *dbcell) Ping() {
	log.Println("cell:", DBCellName, "ping....")
}

func NewDBCell(size uint) types.Cell {
	return &dbcell{BaseCell: BaseCell{name: DBCellName, size: size}}
}
