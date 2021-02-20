package auth

import (
	"net/http"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/handler/api/render"
	"github.com/weekndCN/rw-app/logger"
)

// HandleRegister register a user
func HandleRegister(auth core.AuthStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Get Auth info from Request
		authinfo := &core.Auth{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		if authinfo.Username == "" || authinfo.Password == "" || authinfo.Email == "" {
			render.BadRequest(w, render.ErrMissingParams)
			logger.FromRequest(r).WithError(render.ErrMissingParams).
				Warnln("request: regitser info missing")
			return
		}

		if err := auth.Create(ctx, authinfo); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Errorln("api: create user failed")
			return
		}

		logger.FromRequest(r).Infoln("api: create user success")
		render.JSON(w, "ok", 200)
	}
}
