package data

import (
	"log"
)

func MACDfilter(set *InmemDataset) *InmemDataset {
	outset := NewInmemDataset()
	outset.KeysIndex = set.KeysIndex
	index := outset.KeysIndex["AI.VOL_518"]
	temdmacd := float64(0)
	var tempins *Instance
	for i := 0; i < set.GetLen(); i++ {
		ins := set.GetIndex(i)
		nowmacd := ins.Features.Get(index)
		if temdmacd != nowmacd {

			if nowmacd == 1 {
				tempins = ins
			}
			if nowmacd == 0 {
				if ins.Features.Get(4) > tempins.Features.Get(4) {
					tempins.Output = float64(1)
				} else {
					tempins.Output = float64(0)
				}

				tempins.Features.Set(1, 0)
				tempins.Features.Set(2, 0)
				tempins.Features.Set(3, 0)
				tempins.Features.Set(4, 0)
				tempins.Features.Set(5, 0)
				outset.AddInstance(tempins)
				//log.Println(ins.Features.Get(outset.KeysIndex["AI.MACD0"]))
				log.Printf("%+v", tempins.Features)
				log.Printf("%+v", tempins.Output)
			}
		}
		temdmacd = nowmacd
	}

	return outset
}
