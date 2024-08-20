package middleware

import (
	"errors"
	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var ErrUnauthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error
		if token == "" || username == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
