package deadline

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func Run() {
	ctx, done := context.WithDeadline(context.Background(), time.Now().Add(30*time.Millisecond))

	defer done()

	compute := func() <-chan data {
		ch := make(chan data)

		go func() {
			defer close(ch)

			deadline, ok := ctx.Deadline()

			if ok && time.Since(deadline.Add(-20*time.Millisecond)) > 0 {
				fmt.Println("Timeout.")
				return
			}
			time.Sleep(20 * time.Millisecond)

			select {
			case <-ctx.Done():
				return
			case ch <- data{"abc"}:
			}
		}()

		return ch
	}

	output := <-compute()
	fmt.Printf("Computed data: %s", output.result)
}
