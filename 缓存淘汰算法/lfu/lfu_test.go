package lfu

import (
	"fmt"
	"testing"
)

func TestLFUCache_Set(t *testing.T) {
	//is := is.New(t)
	//cache := NewLFUCache(24, nil)
	//cache.Set("k1", 1)
	//v := cache.Get("k1")
	//is.Equal(v, 1)
	//cache.Del("k1")
	//is.Equal(0, cache.Len())
	keys := make([]string, 0, 8)
	onEvicted := func(k string, v interface{}) {
		keys = append(keys, k)
	}
	cache := NewLFUCache(38, onEvicted)
	cache.Set("key1", "1")
	cache.Set("key2", "2")
	//cache.Get("key1")
	//cache.Get("key1")
	//cache.Get("key2")
	fmt.Println(keys)
	cache.Set("key3", "3")
	cache.Set("key4", "4")
	fmt.Println(keys, cache.Len())
	//expected := []string{"key1", "key3"}
	//is.Equal(keys, expected)
	//is.Equal(2, cache.Len())
}
