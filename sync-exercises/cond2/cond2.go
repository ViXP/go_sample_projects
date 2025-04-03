package cond2

import (
	"fmt"
	"sync"
)

func Run() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	sharedRsc := make(map[string]interface{}, 2)
	c := sync.NewCond(&mx)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) < 1 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) < 2 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
	}()

	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
}
