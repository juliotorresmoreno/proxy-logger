package proxyroutes

import (
	"io"
	"net"
	"net/http"
	"os"
)

type conn struct {
	net.Conn
}

func (el conn) Write(b []byte) (int, error) {
	if os.Getenv("ENVIRONMENT") == "development" {
		os.Stdout.Write(b)
		os.Stdout.Write([]byte("\n\n"))
	}
	return el.Conn.Write(b)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
