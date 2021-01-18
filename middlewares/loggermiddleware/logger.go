package loggermiddleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/juliotorresmoreno/proxy-logger/services/loggerservice"
)

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		httpWriter := loggerservice.NewHTTPWriter(w, r)
		defer func() {
			end := time.Now()
			diff := end.Sub(start).String()
			remoteAddr := strings.Split(r.RemoteAddr, ":")[0]
			statusCode := httpWriter.StatusCode
			log.Printf("%s %v %s %s %v\n", remoteAddr, statusCode, r.Method, r.URL, diff)
		}()
		handler.ServeHTTP(httpWriter, r)

	})
}
