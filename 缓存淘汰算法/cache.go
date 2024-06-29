package cache

import (
	"fmt"
	"runtime"
)

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Del(key string)
	DelOldest()
	Len() int
}
type Value interface {
	Len() int
}

func CalcLen(value interface{}) (n int) {
	switch v := value.(type) {
	case Value:
		n = v.Len()
	case string:
		if runtime.GOARCH == "amd64" {
			n = 16 + len(v)
		} else {
			n = 8 + len(v)
		}
	case bool, uint8, int8:
		n = 1
	case int16, uint16:
		n = 2
	case int32, uint32, float32:
		n = 4
	case int64, uint64, float64:
		n = 8
	case int, uint:
		if runtime.GOARCH == "amd64" {
			n = 8
		} else {
			n = 4
		}
	case complex64:
		n = 8
	case complex128:
		n = 16
	default:
		panic(fmt.Sprintf("%t is not implement cache.value", value))
	}
	return n
}
