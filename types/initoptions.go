package types

import (
	"runtime"
)

var (
	defaultMasterBufferLength = uint(10000)
)

type SkyNetInitOptions struct {
	MasterBufferLength uint
	WorkerSize         int
	IsDebug            bool
}

func (options *SkyNetInitOptions) Init() {
	if options.MasterBufferLength == 0 {
		options.MasterBufferLength = defaultMasterBufferLength
	}

	if options.WorkerSize == 0 {
		options.WorkerSize = runtime.NumCPU()
	}

}
