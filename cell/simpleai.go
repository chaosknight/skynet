package cell

import (
	"log"

	"github.com/chaosknight/skynet/data"
	"github.com/chaosknight/skynet/perceptron"
	"github.com/chaosknight/skynet/types"
	"github.com/chaosknight/skynet/util"
)

const CellAIName = "cellai"

type simpleai struct {
	BaseCell
	perceptron  *perceptron.Perceptron
	initialized bool
}

func (db *simpleai) Init(skynet types.SkyNetInterface) {
	if db.initialized == true {
		log.Fatal("simpleai不能初始化两次")
	}
	db.BaseCell.Init(skynet, db)
	db.initialized = true
	db.perceptron = perceptron.MakePerceptron(10000)
}

func (db *simpleai) Train(set data.InmemDataset) {
	db.perceptron.Train(set)

}

func (db *simpleai) ForwardPass(x *util.Vector) (sum float64) {
	log.Println("x:", x)
	log.Println("r:", db.perceptron.ForwardPass(x) > 0.5)
	return db.perceptron.ForwardPass(x)
}

func (db *simpleai) test() {}

func NewAICell(name string, size uint) types.Cell {
	return &simpleai{BaseCell: BaseCell{name: name, size: size}}
}
