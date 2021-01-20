package proxyroutes

import (
	"io"
	"net/http"
)

func handleHTTP(w http.ResponseWriter, req *reverseRequest) {
	request := *req.Request
	request.URL.Host = req.reverseHOST
	request.Header.Del("Proxy-Authorization")
	request.Header.Del("Proxy-Connection")
	resp, err := http.DefaultTransport.RoundTrip(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
