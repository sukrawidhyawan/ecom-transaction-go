package controllers

import (
	"fmt"
	"net/http"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "order handler")
}
