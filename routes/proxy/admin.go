package proxy

import (
	"bytes"
	"fmt"
	"net/http"
)

//HTTPWriter .
type HTTPWriter struct {
	protocol       string
	statusCode     int
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
	return c
}

//Header .
func (w *HTTPWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *HTTPWriter) Write(b []byte) (int, error) {
	w.responseWriter.Write(b)
	return w.responseWriter.Write(b)
}

//WriteHeader .
func (w *HTTPWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)
}

func (w *HTTPWriter) getContenFromRequest() string {
	req := w.request
	content := fmt.Sprintf("%v %v://%v %v", req.Method, w.protocol, req.Host, req.Proto)
	for header := range req.Header {
		content += fmt.Sprintf("\r\n%v: %v", header, req.Header.Get(header))
	}
	content += fmt.Sprint("\r\n")
	return content
}

//Register .
func (w *HTTPWriter) Register() {
	content := w.getContenFromRequest()
	fmt.Println(content)
}

func registerEvent(w http.ResponseWriter, r *http.Request) *HTTPWriter {
	return NewHTTPWriter(w, r)
}
