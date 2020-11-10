package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/codingno/go-microservices/data"
	"github.com/gorilla/mux"
)

// Products get data products
type Products struct {
	l *log.Logger
}

// NewProducts is returning type of Products
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts is returning all of Products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		// http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		http.Error(rw, "Unable to encoder json", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

// AddProduct is add Product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// UpdateProducts is making changing into  Product
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to decoder json", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
}

// KeyProduct ...
type KeyProduct struct{}

// MiddlewareProductValidation ...
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to decoder json", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
