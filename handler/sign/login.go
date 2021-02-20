package auth

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/core"
)

// HandleLogin handle login fun
func HandleLogin(auth core.AuthStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Get Auth info from Request
		authinfo := &core.Auth{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}

		if authinfo.Email == "" && authinfo.Username == "" {
			logrus.Debugln("Using Email or Username to Login")
			return
		}

		if authinfo.Password == "" {
			logrus.Debugln("Password is needed")
			return
		}

		if authinfo.Username != "" {
			err := auth.Login(ctx, authinfo.Username, authinfo.Password)
			if err != nil {
				logrus.Debugln("Incorrect Auth form data")
				return
			}
			return
		}

		err := auth.Login(ctx, authinfo.Email, authinfo.Password)
		if err != nil {
			logrus.Debugln("Incorrect Auth form data")
			return
		}

		logger := logrus.WithField("login", authinfo.Username)

		logger.Debugf("authentication successful")
	}
}
