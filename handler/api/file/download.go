package file

import (
	"net/http"
	"path"
	"strings"
)

// HandleDownload download file from static files directory
func HandleDownload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get path file
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/octet-stream")
		name := strings.TrimPrefix(r.URL.Path, "/download/")
		w.Header().Set("Content-Disposition: attachment; filename=", name)
		// request file full path
		fullPath := path.Join("uploads", name)
		http.ServeFile(w, r, fullPath)
	}
}
