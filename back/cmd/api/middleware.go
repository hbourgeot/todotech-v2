package main

import (
	"fmt"
	"net/http"
)

func (st *store) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				st.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (st *store) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		apiKey := r.Header.Get("APIKey")

		if apiKey != st.config.apiKey {
			st.invalidKey(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
