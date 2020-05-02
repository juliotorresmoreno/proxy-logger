package status

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Opts struct {
	URL     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]interface{} `json:"headers"`
	Body    interface{}            `json:"body"`
}

func handlePostRequest(c echo.Context) error {
	opts := &Opts{}
	err := c.Bind(opts)
	if err != nil {
		return echo.NewHTTPError(406, err)
	}
	var data string
	if opts.Body != nil {
		data = fmt.Sprintf("%v", opts.Body)
	}
	bodyBuff := bytes.NewBufferString(data)
	req, err := http.NewRequest(opts.Method, opts.URL, bodyBuff)
	if err != nil {
		return echo.NewHTTPError(406, err)
	}
	for key, value := range opts.Headers {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return echo.NewHTTPError(406, err)
	}
	return c.String(204, "")
}

func handleGetClear(c echo.Context) error {
	requests = make([]*Request, 0)
	save()
	return c.JSON(200, requests)
}

type truncateRequest struct {
	ID         string
	StatusCode int
	URI        string
	Method     string
	Time       string
}

func handleGetLogs(c echo.Context) error {
	id := c.Param("id")
	msgRequestID <- id
	return c.String(204, "")
}
