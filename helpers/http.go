package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleStripPrefix(mux *mux.Router, prefix string, handler http.Handler) *mux.Route {
	return mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, handler))
}

func ReadData(reader io.Reader, v interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func HandleHTTPError(w http.ResponseWriter, r *http.Request, err error, status int) bool {
	if err == nil {
		return false
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": err.Error(),
	})
	return true
}
