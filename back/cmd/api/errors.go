package main

import (
	"fmt"
	"net/http"
	"strings"
	"todotech.henrry.online/internal/response"
)

func (st *store) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"error": message}, headers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (st *store) serverError(w http.ResponseWriter, r *http.Request, err error) {

	message := "Internal Server Error"
	st.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (st *store) notFound(w http.ResponseWriter, r *http.Request) {
	message := "Not found"
	st.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (st *store) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	st.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (st *store) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	st.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (st *store) invalidKey(w http.ResponseWriter, r *http.Request) {
	st.errorMessage(w, r, http.StatusUnauthorized, "Invalid API Key", nil)
}
