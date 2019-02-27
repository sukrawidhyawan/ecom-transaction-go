package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/tripLo-team/ecom-transaction-go/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", healthCheckHandler).Methods("GET")
	router.HandleFunc("/api/v1/carts", controllers.CartListItems).Methods("GET")
	router.HandleFunc("/api/v1/carts", controllers.CartAddItem).Methods("POST")
	router.HandleFunc("/api/v1/carts/{userID}/{productID}", controllers.CartGetItem).Methods("GET")
	router.HandleFunc("/api/v1/carts/{userID}/{productID}", controllers.CartPatchItem).Methods("PATCH")
	router.HandleFunc("/api/v1/carts/{userID}/{productID}", controllers.CartDeleteItem).Methods("DELETE")

	router.HandleFunc("/api/v1/orders", controllers.OrderListItems).Methods("GET")
	router.HandleFunc("/api/v1/orders", controllers.OrderAddItem).Methods("POST")
	router.HandleFunc("/api/v1/orders/user/{userID}", controllers.OrderGetUserItems).Methods("GET")
	router.HandleFunc("/api/v1/orders/product/{productID}", controllers.OrderGetProductItems).Methods("GET")
	// router.HandleFunc("/api/v1/orders/:id", controllers.OrderPatchItem).Methods("PATCH")
	// router.HandleFunc("/api/v1/orders/:id", controllers.OrderDeleteItem).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("port " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health Check")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
