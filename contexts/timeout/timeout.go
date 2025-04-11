package timeout

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Run() {
	req, err := http.NewRequest("GET", "https://google.com", nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx, done := context.WithTimeout(req.Context(), 350*time.Millisecond)
	req = req.WithContext(ctx)

	defer done()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
