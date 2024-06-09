package main

import (
	"fmt"
)

type abstractMemory interface {
	storage()
}
type abstractCache interface {
	display()
}
type abstractCpu interface {
	calculate()
}
type absComputer interface {
	createMemory() abstractMemory
	createCache() abstractCache
	createCpu() abstractCpu
}

type IntelMemory struct{}

func (i IntelMemory) storage() {
	fmt.Println("Intel Memory storage")
}

type IntelCache struct{}

func (i IntelCache) display() {
	fmt.Println("Intel Cache display")
}

type IntelCpu struct{}

func (i IntelCpu) calculate() {
	fmt.Println("Intel Cpu calculate")
}

type IntelFactory struct{}

func (i *IntelFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(IntelMemory)
	return memory
}
func (i *IntelFactory) createCache() abstractCache {
	var cache abstractCache
	cache = new(IntelCache)
	return cache
}
func (i IntelFactory) createCpu() abstractCpu {
	var cpu abstractCpu
	cpu = new(IntelCpu)
	return cpu
}

type navidaMemory struct {
}

func (i *navidaMemory) storage() {
	fmt.Println("navida memory storage")
}

type navidaCache struct{}

func (i *navidaCache) display() {
	fmt.Println("navida cache display")
}

type navidaCpu struct{}

func (i *navidaCpu) calculate() {
	fmt.Println("navida cpu calculate")
}

type navidaFactory struct{}

func (n *navidaFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(navidaMemory)
	return memory
}
func (n *navidaFactory) createCache() abstractCache {
	var cache abstractCache
	cache = new(navidaCache)
	return cache
}
func (n *navidaFactory) createCpu() abstractCpu {
	var cpu abstractCpu
	cpu = new(navidaCpu)
	return cpu
}

type kinstonMemory struct {
}

func (k *kinstonMemory) storage() {
	fmt.Println("kinston memory storage")
}

type kinstonCache struct{}

func (k *kinstonCache) display() {
	fmt.Println("kinston cache display")
}

type kinstonCpu struct{}

func (k *kinstonCpu) calculate() {
	fmt.Println("kinston cpu calculate")
}

type KinstonFactory struct{}

func (k *KinstonFactory) createMemory() abstractMemory {
	var memory abstractMemory
	memory = new(kinstonMemory)
	return memory
}
func (k *KinstonFactory) createCpu() abstractCpu {
	var cpu abstractCpu
	cpu = new(kinstonCpu)
	return cpu
}
func (k *KinstonFactory) createCache() abstractCache {
	var cache abstractCache
	cache = new(kinstonCache)
	return cache
}
func main() {
	var iF absComputer
	iF = new(IntelFactory)
	var imem abstractMemory
	imem = iF.createMemory()
	imem.storage()
	icache := iF.createCache()
	icache.display()
	icpu := iF.createCpu()
	icpu.calculate()
	fmt.Println("-----------------------")
	var mixFactory absComputer
	mixFactory = new(IntelFactory)
	mcpu := mixFactory.createCpu()
	mcpu.calculate()
	mixFactory = new(navidaFactory)
	mmem := mixFactory.createMemory()
	mmem.storage()
	mixFactory = new(KinstonFactory)
	mcache := mixFactory.createCache()
	mcache.display()

}
