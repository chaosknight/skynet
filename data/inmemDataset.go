package data

import (
	"log"
)

type InmemDataset struct {
	instances []*Instance

	// 是否所有数据样本都已经添加完毕
	finalized       bool
	FeatureIsSparse bool
	KeysIndex       map[string]int
	//特征向量维度
	FeatureDimension int
}

func NewInmemDataset() *InmemDataset {
	set := new(InmemDataset)
	set.KeysIndex = make(map[string]int)
	return set
}

// 向数据集中添加一个样本
// 成功添加则返回true，否则返回false
func (set *InmemDataset) AddInstance(instance *Instance) bool {
	set.CheckFinalized(false)

	// 添加第一条样本时确定数据集的一些性质
	if len(set.instances) == 0 {
		if instance.Features.IsSparse() {
			set.FeatureIsSparse = true
			set.FeatureDimension = 0
		} else {
			set.FeatureIsSparse = false
			set.FeatureDimension = len(instance.Features.Keys())
		}

	} else {
		if set.FeatureIsSparse {
			if !instance.Features.IsSparse() {
				log.Print("数据集使用稀疏特征而添加的样本不稀疏")
				return false
			}
		} else {
			if instance.Features.IsSparse() {
				log.Print("数据集使用稠密特征而添加的样本稀疏")
				return false
			}

			if set.FeatureDimension != len(instance.Features.Keys()) {
				log.Print("数据集特征数和添加样本的特征数不同")
				return false
			}
		}
	}

	set.instances = append(set.instances, instance)
	return true
}

func (set *InmemDataset) AddSet(that *InmemDataset) {
	for i := 0; i < that.GetLen(); i++ {
		ins := that.GetIndex(i)
		set.AddInstance(ins)
	}
}

func (set *InmemDataset) SetKeys(key string, index int) {
	set.KeysIndex[key] = index
}

func (set *InmemDataset) Finalize() {
	set.CheckFinalized(false)
	set.finalized = true
}

func (set *InmemDataset) GetLen() int {
	return len(set.instances)
}

func (set *InmemDataset) GetIndex(index int) *Instance {
	return set.instances[index]
}

func (set *InmemDataset) CheckFinalized(stat bool) {
	if set.finalized != stat {
		if stat {
			log.Fatal("在遍历数据前必须调用Finalize函数冻结数据")
		} else {
			log.Fatal("冻结数据后不能再对数据集进行修改")
		}
	}
}
