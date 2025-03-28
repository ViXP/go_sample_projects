package example_three

import "fmt"

func Run() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	relayed := <-ch2

	fmt.Printf("Relayed message: %v", relayed)
}

func genMsg(chn chan<- string) {
	chn <- "Some message"
	defer close(chn)
}

func relayMsg(msgCh <-chan string, chn chan<- string) {
	message := <-msgCh
	chn <- message

	defer close(chn)
}
