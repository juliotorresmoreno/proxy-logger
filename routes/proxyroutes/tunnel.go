package proxyroutes

import (
	"net"
	"net/http"
	"time"
)

func handleTunneling(w http.ResponseWriter, r *reverseRequest) {
	// fmt.Println(r.Host)
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	vDestConn := &conn{destConn}
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	vClientConn := &conn{clientConn}
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(vDestConn, vClientConn)
	go transfer(vClientConn, vDestConn)
}
