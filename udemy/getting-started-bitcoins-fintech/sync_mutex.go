package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
	// c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
	// c.mux.Unlock()
}

// channelを使って上手く出来ない場合sync.Mutexを使う
func main() {
	// c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"]++
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			// c["key"]++
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("Key"))
}
