package server

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/juliotorresmoreno/proxy-logger/status"
)

const sep = "\r\n"

type request struct {
	headers   []byte
	body      []byte
	url       string
	method    string
	isTCP     bool
	isCONNECT bool
	upgrade   []byte
}

func parseHTTP(b []byte) *request {
	data := string(b)
	isTCP := isTCP(b)
	binaryData := b
	sepLength := len(sep)

	pos := strings.Index(data, sep)
	connect := data[:pos]
	tmp := strings.Split(connect, " ")
	method := tmp[0]
	url := tmp[1]

	binaryData = binaryData[pos+sepLength:]
	data = data[pos+sepLength:]
	pos = strings.Index(data, sep+sep)

	headers := binaryData[:pos]
	body := binaryData[pos+sepLength*2:]
	isCONNECT := method == "CONNECT"

	return &request{
		method:    method,
		url:       url,
		headers:   headers,
		body:      body,
		isTCP:     isTCP,
		isCONNECT: isCONNECT,
		upgrade:   b,
	}
}

func (svr Server) parseURL(r *request) {
	if strings.Contains(r.url, "http") {
		return
	}
	if svr.Algorithm == "roundrobin" {

	}

	// random
	tmp := svr.getHost(false)
	r.url = tmp + r.url
	host := tmp[strings.Index(tmp, "//")+2:]
	headers := string(r.headers)
	headers = headers + sep + "Host: " + host
	r.headers = []byte(headers)
}

func parseHeaders(req *http.Request, h []byte) {
	s := string(h)
	sr := strings.Split(s, sep)
	for _, line := range sr[1:] {
		t := strings.Split(line, ": ")
		req.Header.Set(t[0], t[1])
	}
}

func (svr Server) call(r *request) (io.Reader, error) {
	svr.parseURL(r)
	fmt.Printf("request: %v %v\n", r.method, r.url)

	buff := bytes.NewBuffer(make([]byte, 0))
	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		return buff, err
	}
	parseHeaders(req, r.headers)
	request, err := http.DefaultClient.Do(req)
	if err != nil {
		return buff, err
	}
	s, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return buff, err
	}

	fmt.Fprintf(buff, "%v %v\n",
		req.Proto,
		request.Status,
	)
	for header := range request.Header {
		fmt.Fprintf(buff, "%v: %v\n", header, request.Header.Get(header))
	}
	fmt.Fprint(buff, sep)
	_, err = buff.Write(s)
	if err != nil {
		return buff, err
	}
	q := bytes.NewBuffer([]byte{})
	fmt.Fprintf(q, "%v %v HTTP/1.1%v", r.method, r.url, sep)
	q.Write(r.headers)
	fmt.Fprintf(q, "%v%v", sep, sep)
	q.Write(r.body)

	log := &status.Request{}

	log.StatusCode = request.StatusCode
	log.Method = r.method
	log.URI = r.url
	log.RawRequest = string(q.Bytes())
	log.RawResponse = string(buff.Bytes())
	log.Raw = string(r.upgrade)
	status.Append(log)

	return buff, nil
}

func (svr Server) handleHTTP(request *request) ([]byte, error) {
	b := make([]byte, 0)
	if request.method == "CONNECT" {
		return b, errors.New("Las peticiones CONNECT no estan soportadas")
	}
	buff, err := svr.call(request)

	if err != nil {
		return b, err
	}
	b, err = ioutil.ReadAll(buff)
	return b, err
}
