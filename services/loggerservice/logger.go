package loggerservice

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

//HTTPWriter .
type HTTPWriter struct {
	Protocol       string
	StatusCode     int
	responseWriter http.ResponseWriter
	request        *http.Request
	buffer         *bytes.Buffer
}

//HTTPResponse .
type HTTPResponse struct {
	http.ResponseWriter
}

//NewHTTPWriter .
func NewHTTPWriter(w http.ResponseWriter, r *http.Request) *HTTPWriter {
	c := new(HTTPWriter)
	c.responseWriter = w
	c.request = r
	c.buffer = bytes.NewBufferString("")
	c.StatusCode = 200
	return c
}

//Header .
func (w *HTTPWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *HTTPWriter) Write(b []byte) (int, error) {
	return w.responseWriter.Write(b)
}

//WriteHeader .
func (w *HTTPWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)
}

func (w *HTTPWriter) getContenFromRequest() string {
	req := w.request
	content := fmt.Sprintf("%v %v://%v %v", req.Method, w.Protocol, req.Host, req.Proto)
	if os.Getenv("ENVIRONMENT") == "development" {
		for header := range req.Header {
			if header == "Authorization" {
				continue
			}
			if header == "Proxy-Authorization" {
				continue
			}
			content += fmt.Sprintf("\r\n%v: %v", header, req.Header.Get(header))
		}
	}
	content += fmt.Sprint("\r\n")
	return content
}

var out io.Writer

// SetLoggerWriter .
func SetLoggerWriter(w io.Writer) {
	out = w
}

//Register .
func (w *HTTPWriter) Register() {
	if out == nil {
		return
	}
	content := w.getContenFromRequest()
	fmt.Fprint(out, content)
}
