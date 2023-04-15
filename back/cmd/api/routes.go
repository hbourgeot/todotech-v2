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
	mux.Handler("POST", "/customers", st.authenticate(http.HandlerFunc(st.AddCustomers)))
	mux.HandlerFunc("GET", "/customers/:id", st.GetCustomers)
	mux.Handler("PUT", "/customers/:id", st.authenticate(http.HandlerFunc(st.UpdateCustomers)))
	mux.Handler("DELETE", "/customers/:id", st.authenticate(http.HandlerFunc(st.DeleteCustomers)))

	// products
	mux.HandlerFunc("GET", "/products", st.GetAllProducts)
	mux.Handler("POST", "/products", st.authenticate(http.HandlerFunc(st.AddProducts)))
	mux.HandlerFunc("GET", "/products/:id", st.GetProducts)
	mux.Handler("PUT", "/products/:id", st.authenticate(http.HandlerFunc(st.UpdateProducts)))
	mux.Handler("DELETE", "/products/:id", st.authenticate(http.HandlerFunc(st.DeleteProducts)))

	// orders
	mux.HandlerFunc("GET", "/orders", st.GetAllOrders)
	mux.Handler("POST", "/orders", st.authenticate(http.HandlerFunc(st.AddOrders)))
	mux.HandlerFunc("GET", "/orders/:id", st.GetOrders)
	mux.Handler("PUT", "/orders/:id", st.authenticate(http.HandlerFunc(st.UpdateOrders)))
	mux.Handler("DELETE", "/orders/:id", st.authenticate(http.HandlerFunc(st.DeleteOrders)))

	return mux
}
