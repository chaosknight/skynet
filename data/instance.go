package data

import (
	"github.com/chaosknight/skynet/util"
)

// 一条数据样本
//
// 对于监督式学习，数据样本通常包含了输入的特征值(features)和输出的目标函数值。
// 对非监督式学习问题，只需要输入的特征值。
type Instance struct {
	// 输入的特征向量
	Features *util.Vector

	// 输出
	// 仅当处理监督式学习问题时需要此项
	// 非监督式学习的数据请使用nil
	Output float64
}
