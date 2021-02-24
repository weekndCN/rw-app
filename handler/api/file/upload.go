package file

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/handler/api/render"
	fapi "github.com/weekndCN/rw-app/store/file"
)

// HandleUpload upload files
func HandleUpload(fs core.FileStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle Option methods
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "authorization")
			return
		}

		ctx := r.Context()
		user := ctx.Value("useranme")

		fmt.Println("user", user)

		r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
		// limit buffer size 32M
		// if overhead then store in disk the rest of data
		r.ParseMultipartForm(30 << 20)

		if r.MultipartForm == nil {
			logrus.Debugln("request MultipartForm is empty")
			render.BadRequest(w, render.ErrNotFound)
			return
		}

		dir := r.MultipartForm.Value["dir"][0]

		files := r.MultipartForm.File["file"]

		for _, file := range files {
			err := fapi.SaveDisk(file, dir)

			if err != nil {
				logrus.Debugln(err)
				render.InternalError(w, err)
				return
			}

			f := &core.File{
				User:       "weeknd",
				Location:   path.Join(dir, file.Filename),
				Name:       file.Filename,
				Size:       file.Size,
				CreateAt:   time.Now(),
				CreateDate: time.Now().String()[0:10],
				Type:       filepath.Ext(file.Filename)[1:],
			}

			err = fs.Create(ctx, f)

			if err != nil {
				logrus.Debugln(err)
				render.InternalError(w, err)
				return
			}
		}

		render.JSON(w, "ok", 200)
	}
}
