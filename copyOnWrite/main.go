package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	batchSize = 1000000
)

type config struct {
	data  map[string]int64
	mutex sync.RWMutex
}

func (c *config) readData(key string) int64 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data[key]
}

func (c *config) setData(data map[string]int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = data
}

func triggerReads(c *config, reply chan int) {
	for i := 0; i <= batchSize; i++ {
		c.readData("naveen")
	}
	reply <- 1
}

func triggerUpdates(c *config, reply chan int) {
	for i := 0; i <= batchSize; i++ {
		c.setData(map[string]int64{
			"naveen": int64(i),
		})
	}
	reply <- 1
}

func withLocks() {
	start := time.Now()
	c := &config{
		data: map[string]int64{
			"naveen": 0,
		},
		mutex: sync.RWMutex{},
	}
	reply := make(chan int, 2)
	go triggerReads(c, reply)
	go triggerUpdates(c, reply)
	<-reply
	<-reply
	fmt.Println("result :", c.readData("naveen"))
	fmt.Println("time taken : ", time.Now().Sub(start))
}

type concurrentConfig struct {
	data unsafe.Pointer
}

func (cc *concurrentConfig) readData(key string) int64 {
	dataPointer := (*map[string]int64)(atomic.LoadPointer(&cc.data))
	return (*dataPointer)[key]
}
func (cc *concurrentConfig) setData(data map[string]int64) {
	atomic.StorePointer(&cc.data, unsafe.Pointer(&data))
}

func triggerReadsOnCOW(c *concurrentConfig, reply chan int) {
	for i := 0; i <= batchSize; i++ {
		c.readData("naveen")
	}
	reply <- 1
}

func triggerUpdatesOnCOW(c *concurrentConfig, reply chan int) {
	for i := 0; i <= batchSize; i++ {
		c.setData(map[string]int64{
			"naveen": int64(i),
		})
	}
	reply <- 1
}

func copyOnWrite() {
	start := time.Now()
	cc := &concurrentConfig{
		data: unsafe.Pointer(&map[string]int64{
			"naveen": 0,
		}),
	}
	reply := make(chan int, 2)
	go triggerReadsOnCOW(cc, reply)
	go triggerUpdatesOnCOW(cc, reply)
	<-reply
	<-reply
	fmt.Println("result :", cc.readData("naveen"))
	fmt.Println("time taken : ", time.Now().Sub(start))
}

func main() {
	//withLocks()
	copyOnWrite()
}
