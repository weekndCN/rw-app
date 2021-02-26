package docker

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/handler/api/render"
	"github.com/weekndCN/rw-app/service/docker"
)

// HandleDockerTail return a docker contailer log tail func
func HandleDockerTail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		fmt.Println("id")
		if id == "" || !ok {
			render.BadRequest(w, errors.New("Must passing a valid containerID"))
			return
		}

		h := w.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Content-Type", "text/event-stream")
		h.Set("Cache-Control", "no-cache")
		h.Set("Connection", "keep-alive")
		h.Set("X-Accel-Buffering", "no")

		f, ok := w.(http.Flusher)
		if !ok {
			return
		}
		// flush header
		io.WriteString(w, ": ping\n\n")
		f.Flush()

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		reader, err := docker.Tail(ctx, id)

		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(w, "event: container-stopped\ndata: end of stream\n\n")
				f.Flush()
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			message := scanner.Text()
			fmt.Fprintf(w, "data: %s\n", message)
			if index := strings.IndexAny(message, " "); index != -1 {
				id := message[:index]
				if _, err := time.Parse(time.RFC3339Nano, id); err == nil {
					fmt.Fprintf(w, "id: %s\n", id)
				}
			}
			fmt.Fprintf(w, "\n")
			f.Flush()
		}

		logrus.Debugf("streaming stopped: %s", id)

		if scanner.Err() == nil {
			logrus.Debugf("container stopped: %v", id)
			fmt.Fprintf(w, "event: container-stopped\ndata: end of stream\n\n")
			f.Flush()
		} else if scanner.Err() != context.Canceled {
			logrus.Errorf("unknown error while streaming %v", scanner.Err())
		}

		logrus.WithField("routines", runtime.NumGoroutine()).Debug("runtime goroutine stats")

		/*

			L:
				for {
					select {
					case <-ctx.Done():
						break L
					case <-time.After(time.Minute * 10):
						defer logReader.Close()
						break L
					default:
						fmt.Printf("%v\n", logReader)
						io.Copy(w, logReader)
						f.Flush()
					}
				}

				io.WriteString(w, "event: error\ndata: eof\n\n")
				f.Flush()
		*/

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
