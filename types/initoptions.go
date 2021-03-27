package types

import (
	"runtime"
)

var (
	defaultMasterBufferLength = uint(100000)
)

type SkyNetInitOptions struct {
	MasterBufferLength uint
	MasterSize         int
}

func (options *SkyNetInitOptions) Init() {
	if options.MasterBufferLength == 0 {
		options.MasterBufferLength = defaultMasterBufferLength
	}

	if options.MasterSize == 0 {
		options.MasterSize = runtime.NumCPU()
	}
}
