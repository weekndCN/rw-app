package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weekndCN/rw-app/core"
	sign "github.com/weekndCN/rw-app/handler/sign"
	"github.com/weekndCN/rw-app/logger"
)

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Auths core.AuthStore
}

// New return a new server
func New(auths core.AuthStore) Server {
	return Server{
		Auths: auths,
	}
}

// Handler return http.Handler
func Handler(s Server) http.Handler {
	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	r.Use(logger.Middleware)
	r.HandleFunc("/login", sign.HandleLogin(s.Auths)).Methods(http.MethodPost)
	r.HandleFunc("/register", sign.HandleRegister(s.Auths)).Methods(http.MethodPost)
	//r.Path("/metrics").Handler(promhttp.Handler())
	//r.Use(auth.HandleAuthMiddler())
	return r
}