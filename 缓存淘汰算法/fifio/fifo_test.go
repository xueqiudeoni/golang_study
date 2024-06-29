package fifo

import (
	"fmt"
	"github.com/matryer/is"
	"testing"
)

func TestSetGet(t *testing.T) {
	is := is.New(t)
	//cache := New(24,nil )
	//cache.DelOldest()
	//cache.Set("k1", 1)
	//v := cache.Get("k1")
	//is.Equal(v, 1)
	//cache.Del("k1")
	//is.Equal(0, cache.Len())

	fmt.Println("--------------------------------------------")
	keys := make([]string, 0)
	onEvicted := func(key string, value interface{}) {
		keys = append(keys, key)
	}
	cache := New(16, onEvicted)

	cache.Set("k1", 1)
	cache.Set("k2", 2)
	cache.Get("k1")
	cache.Set("k3", 3)
	cache.Get("k1")
	cache.Set("k4", 4)
	cache.Del("k4")
	expected := []string{"k1", "k2", "k4"}

	is.Equal(expected, keys)
	is.Equal(1, cache.Len())
}
