package race

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	var t *time.Timer
	start := time.Now()
	ch := make(chan bool)
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Since(start))
		ch <- true
	})
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
