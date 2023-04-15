package main

import (
	"context"
	"net/http"
	"todotech.henrry.online/internal/database/model"
)

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error
type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var RequestValidator RequestValidatorFunc

var ContextInitializer ContextInitializerFunc

func (st *store) ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

func (st *store) initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}

	return ctx
}
