package cache

import (
	"fmt"
	"sync"
	"time"
)

type memCache struct {
	maxMemorySize int64
	// 最大内存字符串表示
	maxMemorySizeStr        string
	usedMemorySize          int64
	values                  map[string]*memCacheValue
	locker                  sync.RWMutex
	clearExpiredKeyInterval time.Duration
}

type memCacheValue struct {
	val interface{}
	// 过期时间
	expireTime time.Time
	// 过期时长，0表示永久有效
	expire int64
	size   int64
}

func NewMemCache() Cache {
	memCache := &memCache{
		values:                  make(map[string]*memCacheValue),
		clearExpiredKeyInterval: time.Second,
	}
	go memCache.clearExpiredKey()
	return memCache
}

// size:1kb 100kb 1MB 1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
	return false
}
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	fmt.Println("set...")
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		size:       GetValSize(val),
	}
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	mc.add(key, v)
	if mc.usedMemorySize > mc.maxMemorySize {
		mc.del(key)
		fmt.Println("maxmemorysize limit, maxmemorysize=%d,usedmemorysize=%d", mc.maxMemorySize, mc.usedMemorySize)
		return false
	}
	return true
}
func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}
func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		delete(mc.values, key)
		mc.usedMemorySize -= tmp.size
	}
}
func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.usedMemorySize += val.size
}

func (mc *memCache) Get(key string) (interface{}, bool) {
	fmt.Println("get...")
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	tmp, ok := mc.get(key)
	if !ok {
		return nil, false
	}
	if tmp.expire != 0 && tmp.expireTime.Before(time.Now()) {
		mc.del(key)
		return nil, false
	}
	return tmp.val, ok
}
func (mc *memCache) Del(key string) bool {
	fmt.Println("del...")
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}
func (mc *memCache) Exist(key string) bool {
	fmt.Println("exist...")
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.get(key)
	return ok
}
func (mc *memCache) Flush() bool {
	fmt.Println("flush...")
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue)
	mc.usedMemorySize = 0
	return true
}
func (mc *memCache) Keys() int64 {
	fmt.Println("keys...")
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}
func (mc *memCache) clearExpiredKey() {
	ticker := time.NewTicker(mc.clearExpiredKeyInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for key, val := range mc.values {
				if val.expire != 0 && val.expireTime.Before(time.Now()) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
