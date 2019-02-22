package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"./controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1", healthCheckHandler).Methods("GET")
	router.HandleFunc("/api/v1/carts", controllers.CartListItems).Methods("GET")
	router.HandleFunc("/api/v1/carts", controllers.CartAddItem).Methods("POST")
	router.HandleFunc("/api/v1/carts/:id", controllers.CartGetItem).Methods("GET")
	router.HandleFunc("/api/v1/carts/:id", controllers.CartPatchItem).Methods("PATCH")
	router.HandleFunc("/api/v1/carts/:id", controllers.CartDeleteItem).Methods("DELETE")

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
