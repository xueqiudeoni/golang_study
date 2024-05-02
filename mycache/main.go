package main

import (
	"mycache/cache"
	"time"
)

func main() {
	cache := cache.NewMemCache()
	cache.SetMaxMemory("200MB")
	cache.Set("int", 1, time.Second*5)
	cache.Set("bool", true, time.Second)
	cache.Get("int")
	cache.Del("int")
	cache.Exist("string")
	cache.Flush()
	cache.Keys()
}
