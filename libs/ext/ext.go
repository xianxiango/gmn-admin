package ext

import (
	"encoding/json"
	"fmt"

	"strconv"
)

func Object2Map(obj interface{}) map[string]interface{} {
	var ret = map[string]interface{}{}
	byt, _ := json.Marshal(obj)
	json.Unmarshal(byt, &ret)
	return ret
}

func Strings2Ints(s []string) []int64 {
	ints := make([]int64, 0, len(s))
	for _, v := range s {
		i, _ := strconv.ParseInt(v, 10, 64)
		ints = append(ints, i)
	}
	return ints
}

func Int64s2Strings(i []int64) []string {
	strs := make([]string, 0, len(i))
	for _, v := range i {
		strs = append(strs, fmt.Sprintf("%d", v))
	}
	return strs
}

func Ints2Strings(i []int) []string {
	strs := make([]string, 0, len(i))
	for _, v := range i {
		strs = append(strs, fmt.Sprintf("%d", v))
	}
	return strs
}
