package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/handler/api/file"
	web "github.com/weekndCN/rw-app/handler/web"
	"github.com/weekndCN/rw-app/logger"
	"github.com/weekndCN/rw-app/token"
)

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Auths core.AuthStore
	Files core.FileStore
}

// New return a new server
func New(auths core.AuthStore, files core.FileStore) Server {
	return Server{
		Auths: auths,
		Files: files,
	}
}

// Handler return http.Handler
func Handler(s Server) http.Handler {
	r := mux.NewRouter()
	r.Use(logger.Middleware)
	r.Use(mux.CORSMethodMiddleware(r))
	// CORSMethodMiddleware add option method on Access-Control-Allow-Methods header flag	r.NewRoute()
	r.HandleFunc("/login", web.HandleLogin(s.Auths)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/register", web.HandleRegister(s.Auths)).Methods(http.MethodPost)
	r.Handle("/", token.Middleware(r))
	r.HandleFunc("/file/upload", file.HandleUpload(s.Files)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/file/list", file.HandleList(s.Files)).Methods(http.MethodGet)
	r.PathPrefix("/download").Handler(file.HandleDownload())
	//r.Path("/metrics").Handler(promhttp.Handler())
	return r
}
