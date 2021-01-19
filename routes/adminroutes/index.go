package adminroutes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func indexFile(w http.ResponseWriter, r *http.Request) {
	indexPath := path.Join("public", "index.html")
	f, err := os.Open(indexPath)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal server error")
	}
	defer f.Close()
	w.WriteHeader(200)
	io.Copy(w, f)
}
