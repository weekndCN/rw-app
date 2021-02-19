package auth

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/core"
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

		if err := auth.Create(ctx, authinfo); err != nil {
			logrus.Debugln("Register failed")
			return
		}

		logrus.Debugln("Register successfully")

	}
}
