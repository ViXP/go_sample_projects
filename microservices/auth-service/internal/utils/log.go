package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func Log(data string) {
	var url string = os.Getenv("LOGGER_URL") + "/api/v1/logs"
	var body []byte = fmt.Appendf(nil, "{\"name\": \"AUTH_REQUEST\", \"data\": \"%v\"}", data)

	http.Post(url, "application/json", bytes.NewBuffer(body))
}
