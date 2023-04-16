package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"

	"todotech.henrry.online/internal/database"
	"todotech.henrry.online/internal/database/model"
	"todotech.henrry.online/internal/request"
	"todotech.henrry.online/internal/response"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (st *store) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	if err := st.ValidateRequest(ctx, r, "customers", model.RetrieveMany); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	records, totalRows, err := database.GetAllCustomers(ctx, "")
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"customers": records,
		"rows":      totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	record, err := database.GetCustomers(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, record)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) AddCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	customers := &model.Customers{}

	if err = request.DecodeJSON(w, r, &customers); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err = customers.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	customers.Prepare()

	if err := customers.Validate(model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	customers, _, err = database.AddCustomers(ctx, customers)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"customer": customers,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Delete); err != nil {
		st.badRequest(w, r, err)
		return
	}

	rowsAffected, err := database.DeleteCustomers(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]int64{"rowsAffected": rowsAffected}
	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) UpdateCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	customers := &model.Customers{}
	if err := request.DecodeJSON(w, r, &customers); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := customers.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	customers.Prepare()

	if err := customers.Validate(model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	customers, _, err = database.UpdateCustomers(ctx, id, customers)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, customers)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveMany); err != nil {
		st.badRequest(w, r, err)
		return
	}

	records, totalRows, err := database.GetAllProducts(ctx, "")
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"products": records,
		"rows":     totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	record, err := database.GetProducts(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, record)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) AddProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	products := &model.Products{}

	if err = request.DecodeJSON(w, r, &products); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err = products.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products.Prepare()

	if err := products.Validate(model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products, _, err = database.AddProducts(ctx, products)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"product": products,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Delete); err != nil {
		st.badRequest(w, r, err)
		return
	}

	rowsAffected, err := database.DeleteProducts(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]int64{"rowsAffected": rowsAffected}
	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	products := &model.Products{}
	if err := request.DecodeJSON(w, r, &products); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := products.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products.Prepare()

	if err := products.Validate(model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products, _, err = database.UpdateProducts(ctx, id, products)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, products)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	if err := st.ValidateRequest(ctx, r, "orders", model.RetrieveMany); err != nil {
		st.badRequest(w, r, err)
		return
	}

	records, totalRows, err := database.GetAllOrders(ctx, "")
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"orders": records,
		"rows":   totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "orders", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	record, err := database.GetOrders(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	product, err := database.GetProducts(ctx, record.IdProduct)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	customer, err := database.GetCustomers(ctx, record.Client)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"order":    record,
		"product":  product,
		"customer": customer,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) AddOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	orders := &model.Orders{}

	if err = request.DecodeJSON(w, r, &orders); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if orders.Date == "" {
		orders.Date = time.Now().Format(time.DateOnly)
	}

	if err = orders.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	orders.Prepare()

	if err := orders.Validate(model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "orders", model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	orders, _, err = database.AddOrders(ctx, orders)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"order": orders,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) DeleteOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Delete); err != nil {
		st.badRequest(w, r, err)
		return
	}

	rowsAffected, err := database.DeleteOrders(ctx, id)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]int64{"rowsAffected": rowsAffected}
	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) UpdateOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	products := &model.Orders{}
	if err := request.DecodeJSON(w, r, &products); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := products.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products.Prepare()

	if err := products.Validate(model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Update); err != nil {
		st.badRequest(w, r, err)
		return
	}

	products, _, err = database.UpdateOrders(ctx, id, products)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, products)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	loginUser := &Login{}
	users := &model.Users{}
	var err error

	if err := request.DecodeJSON(w, r, &loginUser); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := users.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	users.Prepare()

	if err := users.Validate(model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err = st.ValidateRequest(ctx, r, "users", model.RetrieveOne); err != nil {
		st.badRequest(w, r, err)
		return
	}

	users.Email = loginUser.Username

	result := map[string]bool{
		"loggued": false,
	}

	record, err := database.GetUsers(ctx, users.Email)
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword(record.HashedPassword, []byte(loginUser.Password)); err != nil {
		response.JSON(w, http.StatusUnauthorized, result)
		return
	}

	result["loggued"] = true
	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}

func (st *store) AddUsers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	loginUser := &Login{}
	users := &model.Users{}

	if err = request.DecodeJSON(w, r, &loginUser); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err = users.BeforeSave(); err != nil {
		st.badRequest(w, r, err)
		return
	}

	users.Prepare()

	if err := users.Validate(model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}

	if err := st.ValidateRequest(ctx, r, "users", model.Create); err != nil {
		st.badRequest(w, r, err)
		return
	}
	users.Email = loginUser.Username
	users.Created = time.Now()
	users.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(loginUser.Password), 12)
	if err != nil {
		st.badRequest(w, r, err)
		return
	}

	users, _, err = database.AddUsers(ctx, users)
	if err != nil {
		st.serverError(w, r, err)
		return
	}

	result := map[string]any{
		"users": users,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.serverError(w, r, err)
		return
	}
}
