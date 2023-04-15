package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"todotech.henrry.online/internal/database"
	"todotech.henrry.online/internal/database/model"
	"todotech.henrry.online/internal/request"
	"todotech.henrry.online/internal/response"
)

func (st *store) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		st.logger.Fatal(err, nil)
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
		st.logger.Warning("error reading customers: %s", err.Error())
		st.errorMessage(w, r, http.StatusBadRequest, "Error reading customers", nil)
		return
	}

	result := map[string]any{
		"customers": records,
		"rows":      totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	record, err := database.GetCustomers(ctx, id)
	if err != nil {
		st.logger.Warning("error reading customer %s", err.Error())
		return
	}

	err = response.JSON(w, http.StatusOK, record)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) AddCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	customers := &model.Customers{}

	if err = request.DecodeJSON(w, r, &customers); err != nil {
		st.logger.Warning("Error decoding json")
		return
	}

	if err = customers.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 88 %s", err.Error())
		return
	}

	customers.Prepare()

	if err := customers.Validate(model.Create); err != nil {
		st.logger.Warning("error unknown at 95 %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	customers, _, err = database.AddCustomers(ctx, customers)
	if err != nil {
		st.logger.Warning("error creating customer %s", err.Error())
		return
	}

	result := map[string]any{
		"customer": customers,
	}

	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Delete); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	rowsAffected, err := database.DeleteCustomers(ctx, id)
	if err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	result := map[string]int64{"rowsAffected": rowsAffected}
	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) UpdateCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	customers := &model.Customers{}
	if err := request.DecodeJSON(w, r, &customers); err != nil {
		st.logger.Warning("error decoding json %s", err.Error())
		return
	}

	if err := customers.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 156", err.Error())
		return
	}

	customers.Prepare()

	if err := customers.Validate(model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	customers, _, err = database.UpdateCustomers(ctx, id, customers)
	if err != nil {
		st.logger.Warning("error updating customer %s", err.Error())
		return
	}

	err = response.JSON(w, http.StatusOK, customers)
}

func (st *store) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveMany); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	records, totalRows, err := database.GetAllProducts(ctx, "")
	if err != nil {
		st.logger.Warning("error reading products: %s", err.Error())
		st.errorMessage(w, r, http.StatusBadRequest, "Error reading products", nil)
		return
	}

	result := map[string]any{
		"products": records,
		"rows":     totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveOne); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	record, err := database.GetProducts(ctx, id)
	if err != nil {
		st.logger.Warning("error reading product %s", err.Error())
		return
	}

	err = response.JSON(w, http.StatusOK, record)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) AddProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	products := &model.Products{}

	if err = request.DecodeJSON(w, r, &products); err != nil {
		st.logger.Warning("Error decoding json %s", err.Error())
		return
	}

	if err = products.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 88 %s", err.Error())
		return
	}

	products.Prepare()

	if err := products.Validate(model.Create); err != nil {
		st.logger.Warning("error unknown at 95 %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.Create); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	products, _, err = database.AddProducts(ctx, products)
	if err != nil {
		st.logger.Warning("error creating customer %s", err.Error())
		return
	}

	result := map[string]any{
		"customer": products,
	}

	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Delete); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	rowsAffected, err := database.DeleteProducts(ctx, id)
	if err != nil {
		st.logger.Warning("error validating request %s", err.Error())
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
		st.logger.Warning("failed atoi")
		return
	}

	products := &model.Products{}
	if err := request.DecodeJSON(w, r, &products); err != nil {
		st.logger.Warning("error decoding json %s", err.Error())
		return
	}

	if err := products.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 156", err.Error())
		return
	}

	products.Prepare()

	if err := products.Validate(model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	products, _, err = database.UpdateProducts(ctx, id, products)
	if err != nil {
		st.logger.Warning("error updating customer %s", err.Error())
		return
	}

	err = response.JSON(w, http.StatusOK, products)
}

func (st *store) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	if err := st.ValidateRequest(ctx, r, "orders", model.RetrieveMany); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	records, totalRows, err := database.GetAllOrders(ctx, "")
	if err != nil {
		st.logger.Warning("error reading orders: %s", err.Error())
		st.errorMessage(w, r, http.StatusBadRequest, "Error reading orders", nil)
		return
	}

	result := map[string]any{
		"products": records,
		"rows":     totalRows,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "orders", model.RetrieveOne); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	record, err := database.GetOrders(ctx, id)
	if err != nil {
		st.logger.Warning("error reading product %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.RetrieveOne); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	product, err := database.GetProducts(ctx, record.IdProduct)
	if err != nil {
		st.logger.Warning("error reading product %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "customers", model.RetrieveOne); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	customer, err := database.GetCustomers(ctx, record.Client)
	if err != nil {
		st.logger.Warning("error reading customer %s", err.Error())
		return
	}

	result := map[string]any{
		"order":    record,
		"product":  product,
		"customer": customer,
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		st.logger.Warning("error serving data %s", err.Error())
	}
}

func (st *store) AddOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)
	var err error
	orders := &model.Orders{}

	if err = request.DecodeJSON(w, r, &orders); err != nil {
		st.logger.Warning("Error decoding json %s", err.Error())
		return
	}

	if orders.Date == "" {
		orders.Date = time.Now().Format(time.DateOnly)
		fmt.Println(orders.Date)
	}

	if err = orders.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 414 %s", err.Error())
		return
	}

	orders.Prepare()

	if err := orders.Validate(model.Create); err != nil {
		st.logger.Warning("error unknown at 421 %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "orders", model.Create); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	orders, _, err = database.AddOrders(ctx, orders)
	if err != nil {
		st.logger.Warning("error creating customer %s", err.Error())
		return
	}

	result := map[string]any{
		"customer": orders,
	}

	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) DeleteOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Delete); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	rowsAffected, err := database.DeleteOrders(ctx, id)
	if err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	result := map[string]int64{"rowsAffected": rowsAffected}
	err = response.JSON(w, http.StatusOK, result)
}

func (st *store) UpdateOrders(w http.ResponseWriter, r *http.Request) {
	ctx := st.initializeContext(r)

	params := httprouter.ParamsFromContext(ctx)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		st.logger.Warning("failed atoi")
		return
	}

	products := &model.Orders{}
	if err := request.DecodeJSON(w, r, &products); err != nil {
		st.logger.Warning("error decoding json %s", err.Error())
		return
	}

	if err := products.BeforeSave(); err != nil {
		st.logger.Warning("error unknown at 156", err.Error())
		return
	}

	products.Prepare()

	if err := products.Validate(model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	if err := st.ValidateRequest(ctx, r, "products", model.Update); err != nil {
		st.logger.Warning("error validating request %s", err.Error())
		return
	}

	products, _, err = database.UpdateOrders(ctx, id, products)
	if err != nil {
		st.logger.Warning("error updating customer %s", err.Error())
		return
	}

	err = response.JSON(w, http.StatusOK, products)
}
