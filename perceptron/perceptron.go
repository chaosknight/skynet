package perceptron

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/chaosknight/skynet/data"
	"github.com/chaosknight/skynet/util"
)

type Perceptron struct {
	weights *util.Vector
	bias    float64
	epochs  int
}

func (a *Perceptron) sigmoid(x float64) float64 { //Sigmoid Activation
	return 1.0 / (1.0 + math.Exp(-x))
}

func (a *Perceptron) ForwardPass(x *util.Vector) (sum float64) { //Forward Propagation
	return a.sigmoid(util.VecDotProduct(a.weights, x) + a.bias)
}

func (a *Perceptron) gradW(x *util.Vector, y float64) *util.Vector { //Calculate Gradients of Weights
	pred := a.ForwardPass(x)
	res := x.Populate()
	res.Increment(x, -(pred-y)*pred*(1-pred))
	return res
}

func (a *Perceptron) gradB(x *util.Vector, y float64) float64 { //Calculate Gradients of Bias
	pred := a.ForwardPass(x)
	return -(pred - y) * pred * (1 - pred)
}

func (a *Perceptron) Train(set data.InmemDataset) { //Train the Perceptron for n epochs
	a.weights = util.NewVector(set.FeatureDimension)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < set.FeatureDimension; i++ {
		a.weights.Set(i, rand.Float64())
	}

	dw := util.NewVector(set.FeatureDimension)
	for i := 0; i < a.epochs; i++ {
		dw.Clear()
		db := 0.0
		for j := 0; i < set.GetLen(); i++ {
			ins := set.GetIndex(j)
			dw.Increment(a.gradW(ins.Features, ins.Output), 1)
			db += a.gradB(ins.Features, ins.Output)
		}

		dw.Increment(dw, 2/float64(set.GetLen()))
		a.weights.Increment(dw, 1)

		a.bias += db * 2 / float64(set.GetLen())
	}
	log.Printf("%+v", a.weights)

}

func MakePerceptron(epochs int) *Perceptron {
	return &Perceptron{
		epochs: epochs, //Number of Epoch
		bias:   0,
	}
}
