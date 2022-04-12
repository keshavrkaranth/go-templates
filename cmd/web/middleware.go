package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//NoSurf is csrf protection to post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves session for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
