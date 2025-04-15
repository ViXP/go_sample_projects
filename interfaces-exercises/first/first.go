package first

import "fmt"

type ByteCounter int

func (bc *ByteCounter) Write(bytes []byte) (int, error) {
	*bc += ByteCounter(len(bytes))
	return int(*bc), nil
}

func Run() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
