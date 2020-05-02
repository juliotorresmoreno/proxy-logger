package server

import (
	"net"

	"github.com/juliotorresmoreno/proxy-logger/status"
)

func (svr Server) handleCONNECT(conn net.Conn, b []byte) error {
	r := parseHTTP(b)
	err := svr.handleTCP(conn, b, true, nil)
	if err == nil {
		log := &status.Request{
			StatusCode:  200,
			Method:      r.method,
			URI:         r.url,
			RawRequest:  string(r.upgrade),
			RawResponse: "El contendo esta cifrado",
			Raw:         string(r.upgrade),
		}
		status.Append(log)
		return nil
	}
	log := &status.Request{
		StatusCode:  500,
		Method:      r.method,
		URI:         r.url,
		RawRequest:  string(r.upgrade),
		RawResponse: "El contendo esta cifrado",
		Raw:         string(r.upgrade),
	}
	status.Append(log)

	return err
}
