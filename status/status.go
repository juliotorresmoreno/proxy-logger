package status

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

type Request struct {
	ID         string
	StatusCode int
	URI        string
	Method     string
	Time       string

	RawRequest  string
	RawResponse string
	Raw         string `json:"-"`
}

type response struct {
	req *http.Request
	res http.ResponseWriter
	id  string
}

var truncateRequests = make([]*truncateRequest, 0)
var requests = make([]*Request, 0)

var msgAppendReq = make(chan *Request)
var msgRequests = make(chan bool)
var msgRequestID = make(chan string)

func init() {
	go channels()
}

func msgAppendReqFn(r *Request) {
	r.ID = bson.NewObjectId().Hex()
	r.Time = time.Now().Format(time.RFC3339)
	requests = append(requests, r)
	length := len(requests)
	if length > 1000 {
		requests = requests[length-1000:]
	}
	truncateRequests = make([]*truncateRequest, 0, len(requests))
	for _, req := range requests {
		tmp := &truncateRequest{
			ID:         req.ID,
			Method:     req.Method,
			StatusCode: req.StatusCode,
			Time:       req.Time,
			URI:        req.URI,
		}
		truncateRequests = append(truncateRequests, tmp)
	}
}

func msgRequestsFn() {
	buff := bytes.NewBufferString("")
	json.NewEncoder(buff).Encode(map[string]interface{}{
		"type": "requests",
		"data": requests,
	})
	Send(buff.String())
	save()
}

func msgRequestIDFn(reqID string) {
	for _, req := range requests {
		if req.ID == reqID {
			buff := bytes.NewBufferString("")
			json.NewEncoder(buff).Encode(map[string]interface{}{
				"type": "request",
				"data": req,
			})
			Send(buff.String())
			break
		}
	}
}

func channels() {
	for {
		select {
		case req := <-msgAppendReq:
			msgAppendReqFn(req)
		case _ = <-msgRequests:
			msgRequestsFn()
		case reqID := <-msgRequestID:
			msgRequestIDFn(reqID)
		}
	}
}

func Append(request *Request) {
	go func() {
		msgAppendReq <- request
		msgRequests <- true
	}()
}

func read() {
	f, err := os.Open("/tmp/requests.json")
	if err == nil {
		json.NewDecoder(f).Decode(&requests)
	}
}

func save() {
	f, err := os.OpenFile("/tmp/requests.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	json.NewEncoder(f).Encode(requests)
}

func NewServer() *echo.Echo {
	read()

	svr := echo.New()
	svr.Static("/static", "bower_components")
	svr.Static("/", "public")
	svr.POST("/request", handlePostRequest)
	svr.GET("/logs/http/:id", handleGetLogs)
	svr.GET("/clear", handleGetClear)
	svr.GET("/ws", handleGetWebsocket)
	return svr
}
