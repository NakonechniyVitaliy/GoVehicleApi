package helper

import "net/url"

const (
	LocalHost = "localhost:8082"
)

var TcUrl = url.URL{
	Scheme: "http",
	Host:   LocalHost,
}
