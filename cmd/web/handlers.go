package main

import (
	"encoding/json"
	"github.com/Bug-daulet/qlt_task/pkg/models"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (a *app) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments := a.dbPoolPayment.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func (a *app) GetPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payment := a.dbPoolPayment.Get(id)
	json.NewEncoder(w).Encode(payment)
}

func (a *app) CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payment models.Payment
	_ = json.NewDecoder(r.Body).Decode(&payment)
	payment.ID = a.dbPoolPayment.Save(&payment)
	json.NewEncoder(w).Encode(payment)
}

func (a *app) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var payment models.Payment
	_ = json.NewDecoder(r.Body).Decode(&payment)
	payment.ID = id
	a.dbPoolPayment.Update(&payment)
	json.NewEncoder(w).Encode(payment)
}

func (a *app) DeletePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	a.dbPoolPayment.Delete(id)
}


func (a *app) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := a.dbPoolCategory.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (a *app) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category := a.dbPoolCategory.Get(id)
	json.NewEncoder(w).Encode(category)
}

func (a *app) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	category.ID = a.dbPoolCategory.Save(&category)
	json.NewEncoder(w).Encode(category)
}

func (a *app) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var category models.Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	category.ID = id
	a.dbPoolCategory.Update(&category)
	json.NewEncoder(w).Encode(category)
}

func (a *app) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	a.dbPoolCategory.Delete(id)
}

