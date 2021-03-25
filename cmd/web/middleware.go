package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Hit the page")
//		next.ServeHTTP(w, r)
//	})
//}

func NoSurf(next http.Handler) http.Handler {
	csrfHandeler := nosurf.New(next)

	csrfHandeler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.Production,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandeler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}