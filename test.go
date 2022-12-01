package main

import (
	"github.com/chaosknight/skynet/cell"
	// "github.com/chaosknight/skynet/types"
)

const cellScalName = "cellscal"

type Scal struct {
	id uint
	cell.Cell
}

func NewScal() *Scal {
	ce := cell.GO(cellScalName, 10)
	return &Scal{
		Cell: *ce,
		id:   0,
	}
}
