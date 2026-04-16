package helper

import (
	"net/url"
	"os"
)

var TcUrl = func() url.URL {
	base := os.Getenv("TEST_BASE_URL")
	if base == "" {
		base = "http://localhost:8082"
	}
	u, _ := url.Parse(base)
	return *u
}()
