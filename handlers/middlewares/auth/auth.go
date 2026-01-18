package auth

import (
	"context"
	"net/http"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/app/models"
	"github.com/SmurfsAtWork/lilpapa/handlers/middlewares"
)

// Context keys
const (
	UserKey = "user"
)

type Middleware struct {
	usecases *actions.Actions
}

// New returns a new auth middleware instance.
func New(usecases *actions.Actions) *Middleware {
	return &Middleware{
		usecases: usecases,
	}
}

func (a *Middleware) AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := a.authenticate(r)
		if err != nil {
			middlewares.HandleErrorResponse(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthApi authenticates an API's handler.
func (a *Middleware) AuthApi(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := a.authenticate(r)
		if err != nil {
			middlewares.HandleErrorResponse(w, err)
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, user)
		h(w, r.WithContext(ctx))
	}
}

// OptionalAuthApi authenticates an API's handler optionally (without 401).
func (a *Middleware) OptionalAuthApi(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := a.authenticate(r)
		if err != nil {
			h(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, user)
		h(w, r.WithContext(ctx))
	}
}

func (a *Middleware) authenticate(r *http.Request) (models.User, error) {
	sessionToken, ok := r.Header["Authorization"]
	if !ok {
		return models.User{}, &actions.ErrInvalidSessionToken{}
	}

	return a.usecases.AuthenticateUser(sessionToken[0])
}
