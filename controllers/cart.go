package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

func CartListItems(w http.ResponseWriter, r *http.Request) {
	data := models.ListAllCart()
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
	// fmt.Fprintf(w, "CartListItems handler")
}

func CartAddItem(w http.ResponseWriter, r *http.Request) {
	cart := &models.Cart{}
	err := json.NewDecoder(r.Body).Decode(cart)

	if err != nil {
		utils.Respond(w, utils.Message(500, "Error while decoding request body"))
		return
	}
	data := cart.Create()
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func CartGetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])
	productID, _ := strconv.Atoi(vars["productID"])
	data := models.GetCart(userID, productID)
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func CartPatchItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])
	productID, _ := strconv.Atoi(vars["productID"])

	cart := &models.Cart{}
	err := json.NewDecoder(r.Body).Decode(cart)

	if err != nil {
		utils.Respond(w, utils.Message(500, "Error while decoding request body"))
		return
	}
	cart.UserID = userID
	cart.ProductID = productID
	data := cart.Update()
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func CartDeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CartDeleteItem handler")
}
