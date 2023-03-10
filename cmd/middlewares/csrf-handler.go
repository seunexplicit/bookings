package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func CSRFHandler(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Name:     "csrf_token",
        Path:     "/",
        HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure: false,
	})

	return csrfHandler
}