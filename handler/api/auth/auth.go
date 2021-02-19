package auth

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/weekndCN/rw-app/handler/api/request"
	"github.com/weekndCN/rw-app/logger"
)

// HandleAuthMiddler retrun an http.HandlerFunc
func HandleAuthMiddler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				log := logger.FromContext(ctx)
				// Get auth
				auth, ok := request.AuthFrom(ctx)

				if auth.Username == "" || auth.Email == "" || !ok {
					next.ServeHTTP(w, r)
					log.Debugln("access denfied")
					return
				}

				if auth.Username != "" {
					log = log.WithFields(logrus.Fields{"User login:": auth.Username})
				}

				if auth.Email != "" {
					log.WithFields(logrus.Fields{"email login:": auth.Email})
				}

				ctx = logger.WithContext(ctx, log)
				next.ServeHTTP(w, r.WithContext(request.WithAuth(ctx, auth)))
			},
		)
	}
}
