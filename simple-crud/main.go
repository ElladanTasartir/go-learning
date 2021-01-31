package main

import (
	"fmt"
	"net/http"
	"simplecrud/products"
)

func main() {
	router := products.Init()
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
