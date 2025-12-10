package http

import (
	"net/http"
)

type Reachable struct {
	Addr     string
	Response *http.Response

	HTTPsEnabled bool
}

func (r *Reachable) String() string {
	if r.Response == nil {
		return "Unreachable through HTTP(" + r.Addr + ":80) and HTTPs(" + r.Addr + ":443)"
	}

	var reachableStr = r.Response.Proto + " " + r.Addr + ": " + r.Response.Status + " "
	if r.Response.StatusCode >= 300 && r.Response.StatusCode < 400 {
		reachableStr = reachableStr + "Redirect to " + r.Response.Header.Get("Location")
	}

	return reachableStr
}

func getRequest(url string) (*http.Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func IsAddrHTTPReachable(addr string) (*Reachable, error) {
	res, err := getRequest("http://" + addr)
	if err != nil {
		return &Reachable{Addr: addr, Response: res}, err
	}

	return &Reachable{Addr: addr, Response: res}, nil
}

func IsAddrHTTPSReachable(addr string) (*Reachable, error) {
	res, err := getRequest("https://" + addr)
	if err != nil {
		return &Reachable{Addr: addr, Response: res}, err
	}

	return &Reachable{Addr: addr, Response: res, HTTPsEnabled: true}, nil
}

func IsAddrReachable(addr string) (*Reachable, error) {
	tlsConn, err := tlsDialer.Dial("tcp", addr+":443")
	if err != nil {
		return IsAddrHTTPReachable(addr)
	}
	defer tlsConn.Close()

	return IsAddrHTTPSReachable(addr)
}

func AreAddrsReachable(addresses []string) ([]*Reachable, error) {
	reachableAddrs := make([]*Reachable, len(addresses))
	for i, addr := range addresses {
		reachableAddrs[i], _ = IsAddrReachable(addr)
	}

	return reachableAddrs, nil
}
