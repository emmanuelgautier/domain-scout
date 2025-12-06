package http

import (
	"net"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

var tlsDialer = &net.Dialer{
	Timeout: 10 * time.Second,
}
