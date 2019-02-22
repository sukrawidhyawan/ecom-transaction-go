package controllers

import (
	"fmt"
	"net/http"
)

func CartListItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartListItems handler")
}

func CartAddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartAddItem handler")
}

func CartGetItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartGetItem handler")
}

func CartPatchItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartPatchItem handler")
}

func CartDeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartDeleteItem handler")
}
