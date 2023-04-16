package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (st *store) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(st.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(st.methodNotAllowed)

	mux.Handler("GET", "/", st.authenticate(http.HandlerFunc(st.status)))

	// customers
	mux.HandlerFunc("GET", "/customers", st.GetAllCustomers)
	mux.HandlerFunc("POST", "/customers", st.AddCustomers)
	mux.HandlerFunc("GET", "/customers/:id", st.GetCustomers)
	mux.HandlerFunc("PUT", "/customers/:id", st.UpdateCustomers)
	mux.HandlerFunc("DELETE", "/customers/:id", st.DeleteCustomers)

	// products
	mux.HandlerFunc("GET", "/products", st.GetAllProducts)
	mux.HandlerFunc("POST", "/products", st.AddProducts)
	mux.HandlerFunc("GET", "/products/:id", st.GetProducts)
	mux.HandlerFunc("PUT", "/products/:id", st.UpdateProducts)
	mux.HandlerFunc("DELETE", "/products/:id", st.DeleteProducts)

	// orders
	mux.HandlerFunc("GET", "/orders", st.GetAllOrders)
	mux.HandlerFunc("POST", "/orders", st.AddOrders)
	mux.HandlerFunc("GET", "/orders/:id", st.GetOrders)
	mux.HandlerFunc("PUT", "/orders/:id", st.UpdateOrders)
	mux.HandlerFunc("DELETE", "/orders/:id", st.DeleteOrders)

	mux.HandlerFunc("POST", "/login", st.AddUsers)
	mux.HandlerFunc("GET", "/login", st.GetUsers)
	return st.recoverPanic(st.authenticate(mux))
}
