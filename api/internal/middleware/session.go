package middleware

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func Session(manager *scs.SessionManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return manager.LoadAndSave(next)
	}
}
