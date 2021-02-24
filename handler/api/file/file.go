package file

import (
	"net/http"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/handler/api/render"
)

// filter .
type filter struct {
	// file name
	name string
	// file type
	fileType string
	// file create date
	createAt string
}

// HandleList handle file
func HandleList(fs core.FileStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add request parameter
		params := r.URL.Query()

		var files *[]core.File
		var err error
		// if url request no parameter then search all data
		if len(params) == 0 {
			files, err = fs.List(ctx)
			if err != nil {
				render.InternalError(w, err)
				return
			}

			render.JSON(w, files, 200)
			return
		}

		f := &core.File{
			Type:       params.Get("type"),
			Name:       params.Get("name"),
			CreateDate: params.Get("time"),
		}

		files, err = fs.Find(ctx, f)

		if err != nil {
			render.InternalError(w, err)
			return
		}

		render.JSON(w, files, 200)
		return
	}
}
