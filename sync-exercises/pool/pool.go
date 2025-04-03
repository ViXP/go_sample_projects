package pool

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new output buffer...")
		return new(bytes.Buffer)
	},
}

func Run() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	wg.Add(2)
	go log(os.Stdout, "log1", &wg, &mx)
	go log(os.Stdout, "log2", &wg, &mx)

	wg.Wait()
}

func log(w io.Writer, debug string, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	mx.Lock()
	b := bufferPool.Get().(*bytes.Buffer)
	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")
	w.Write(b.Bytes())
	b.Reset()
	bufferPool.Put(b)
	mx.Unlock()
}
