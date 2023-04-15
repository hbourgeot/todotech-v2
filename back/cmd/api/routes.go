package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (st *store) routes() http.Handler {
	mux := httprouter.New()

	mux.HandlerFunc("GET", "/", st.status)
	//mux.NotFound = http.HandlerFunc(app.notFound)
	//mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	//mux.HandlerFunc("GET", "/status", app.status)
	//mux.HandlerFunc("POST", "/users", app.createUser)
	//mux.HandlerFunc("POST", "/authentication-tokens", app.createAuthenticationToken)

	//mux.Handler("GET", "/protected", app.requireAuthenticatedUser(http.HandlerFunc(app.protected)))

	//mux.Handler("GET", "/basic-auth-protected", app.requireBasicAuthentication(http.HandlerFunc(app.protected)))

	return mux
}
