package proxyroutes

import (
	"net/http"

	"github.com/juliotorresmoreno/proxy-logger/config"
)

type reverseRequest struct {
	*http.Request
	reverseHOST string
}

func reverseURI(r *http.Request) *reverseRequest {
	config, err := config.GetConfig()
	if err == nil {
		rv := config.GetReverse()
		if reverse, ok := rv[r.Host]; ok {
			return &reverseRequest{r, reverse}
		}
	}
	return &reverseRequest{r, r.Host}
}
