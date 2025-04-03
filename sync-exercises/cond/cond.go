package cond

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func Run() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	conditional := sync.NewCond(&mx)

	wg.Add(1)

	go func() {
		defer wg.Done()

		conditional.L.Lock()
		for len(sharedRsc) == 0 {
			conditional.Wait()
		}
		conditional.L.Unlock()

		fmt.Printf("The resource is: %v", sharedRsc["rsc1"])
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		conditional.L.Lock()
		sharedRsc["rsc1"] = 10
		conditional.Signal()
		conditional.L.Unlock()
	}()

	wg.Wait()
}
