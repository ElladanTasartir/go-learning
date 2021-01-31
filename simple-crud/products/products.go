package products

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var products = []Product{}
var idCount int

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	return router
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read request")
	}

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read request")
	}

	var product Product
	product.ID = idCount + 1
	idCount++
	json.Unmarshal(body, &product)
	products = append(products, product)

	w.WriteHeader(201)

	json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read request")
	}

	var sentProduct Product
	json.Unmarshal(body, &sentProduct)
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Param sent in wrong format")
	}
	for index, product := range products {
		if product.ID == id {
			if sentProduct.Name != "" {
				products[index].Name = sentProduct.Name
			}
			if sentProduct.Price != 0 {
				products[index].Price = sentProduct.Price
			}

			json.NewEncoder(w).Encode(products[index])
		}
	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]

	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Param sent in wrong format")
	}

	for index, product := range products {
		if product.ID == id {
			products = append(products[:index], products[index+1:]...)

			w.WriteHeader(204)
			json.NewEncoder(w).Encode(nil)
		}
	}
}
