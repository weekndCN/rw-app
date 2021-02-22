package file

import (
	"net/http"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/handler/api/render"
)

// HandleList handle file
func HandleList(fs core.FileStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		files, err := fs.List(ctx)
		if err != nil {
			render.InternalError(w, err)
			return
		}

		render.JSON(w, files, 200)
	}
}
