package contrib

import (
	"io/ioutil"
	"log"
	"skynet/data"
	"skynet/util"
	"strconv"
	"strings"
)

func LoadLibSVMDataset(path string, usingSparseRepresentation bool) *data.InmemDataset {
	log.Print("载入libsvm格式文件", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("无法打开文件\"%v\"，错误提示：%v\n", path, err)
	}
	lines := strings.Split(string(content), "\n")
	set := data.NewInmemDataset()
	for _, l := range lines {
		if l == "" {
			continue
		}
		fields := strings.Split(strings.TrimSpace(l), "	")

		if len(fields) <= 6 {

			continue
		}
		if len(set.KeysIndex) == 0 {
			for i, v := range fields {
				switch i {
				case 0:
					set.SetKeys("time", i+1)
				case 1:
					set.SetKeys("O", i+1)
				case 2:
					set.SetKeys("H", i+1)
				case 3:
					set.SetKeys("L", i+1)
				case 4:
					set.SetKeys("C", i+1)
				case 5:
					set.SetKeys("vol", i+1)
				default:
					set.SetKeys(strings.TrimSpace(v), i+1)
				}
			}
			continue
		}
		instance := new(data.Instance)

		instance.Features = util.NewVector(len(set.KeysIndex) + 1)
		instance.Features.Set(0, 0)
		var tem = ""
		for i := 1; i < len(fields); i++ {
			tem = strings.TrimSpace(fields[i])
			if tem == "" {
				continue
			}
			value, _ := strconv.ParseFloat(tem, 64)
			instance.Features.Set(i, value)

		}
		set.AddInstance(instance)
		// log.Printf("%+v", instance.Features)
	}

	return set
}
