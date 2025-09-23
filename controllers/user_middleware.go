package controllers

import (
	"net/http"

	"github.com/fadyboy/lenslocked/context"
	"github.com/fadyboy/lenslocked/models"
)

type UserMiddleWare struct {
	SessionService *models.SessionService
}

func (umw UserMiddleWare) SetUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read cookie here to get token
		token, err := readCookie(r, CookieSession)
		if err != nil {
			// if no token found, proceed without setting user
			next.ServeHTTP(w, r)
			return
		}

		// look up user using token
		user, err := umw.SessionService.User(token)
		if err != nil {
			// invalid/expired token so no user returned to set user
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		ctx = context.WithUser(ctx, user)

		// get new request with context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}


func (umw UserMiddleWare) RequireContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := context.User(r.Context())
		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
