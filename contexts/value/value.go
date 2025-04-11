package value

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKeyType string

var db = database{
	"jane": true,
}

func Run() {
	ctx, done := context.WithCancel(context.Background())

	defer done()

	process(ctx, "jane")
	process(ctx, "john")
}

func process(ctx context.Context, id string) {
	ctx = context.WithValue(ctx, userIDKeyType("id"), id)

	select {
	case <-ctx.Done():
		return
	case status := <-checkMembership(ctx):
		fmt.Printf("\nmembership status of the reader : %s : %v\n", id, status)
	}
}

func checkMembership(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		ch <- db[ctx.Value(userIDKeyType("id")).(string)]
	}()
	return ch
}
