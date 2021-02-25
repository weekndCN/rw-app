package docker

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/weekndCN/rw-app/handler/api/render"
	"github.com/weekndCN/rw-app/service/docker"
)

// HandleDockerTail return a docker contailer log tail func
func HandleDockerTail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Content-Type", "text/event-stream")
		h.Set("Cache-Control", "no-cache")
		h.Set("Connection", "keep-alive")
		h.Set("X-Accel-Buffering", "no")

		f, ok := w.(http.Flusher)
		if !ok {
			return
		}

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		logReader, err := docker.Tail(ctx, "28cb50796bce")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
	L:
		for {
			select {
			case <-ctx.Done():
				break L
			case <-time.After(time.Minute):
				break L
			default:
				io.Copy(w, logReader)
				f.Flush()
			}
		}
		io.WriteString(w, "event: error\ndata: eof\n\n")
		f.Flush()
	}
}

// HandleDockerList  list container
func HandleDockerList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// cs receiver containers
		cs, err := docker.List()
		if err != nil {
			render.InternalError(w, err)
			return
		}

		render.JSON(w, cs, 200)
		return
	}
}

// HandleDockerStart start container
func HandleDockerStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if id == "" || !ok {
			render.BadRequest(w, errors.New("Must passing a valid containerID"))
			return
		}

		err := docker.Start(id)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, "ok", 200)
		return
	}
}

// HandleDockerStop stop container
func HandleDockerStop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if id == "" || !ok {
			render.BadRequest(w, errors.New("Must passing a valid containerID"))
			return
		}

		err := docker.Stop(id)
		if err != nil {
			render.InternalError(w, err)
			return
		}
		render.JSON(w, "ok", 200)
		return
	}
}
