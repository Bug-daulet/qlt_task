package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (a *app) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/payments", a.GetAllPayments)
	mux.Get("/payments/{id:[0-9]+}", a.GetPayment)
	mux.Post("/payments/create", a.CreatePayment)
	mux.Post("/payments/edit/{id:[0-9]+}", a.UpdatePayment)
	mux.Post("/payments/delete/{id:[0-9]+}", a.DeletePayment)

	mux.Get("/categories", a.GetAllCategories)
	mux.Get("/categories/{id:[0-9]+}", a.GetCategory)
	mux.Post("/categories/create", a.CreateCategory)
	mux.Post("/categories/edit/{id:[0-9]+}", a.UpdateCategory)
	mux.Post("/categories/delete/{id:[0-9]+}", a.DeleteCategory)

	return mux
}

