package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tripLo-team/ecom-transaction-go/utils"
	"github.com/gorilla/mux"
	"github.com/tripLo-team/ecom-transaction-go/models"
)

func OrderListItems(w http.ResponseWriter, r *http.Request) {
	data := models.ListAllOrder()
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func OrderAddItem(w http.ResponseWriter, r *http.Request) {
	type OrderRequest struct {
		Orders models.Orders `json:"orders"`
	}
	request := &OrderRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		utils.Respond(w, utils.Message(500, "Error while decoding request body"))
		return
	}
	data := (request.Orders).CreateBulk()
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func OrderGetUserItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["userID"])
	fmt.Println("userId " + strconv.Itoa(userID))

	data := models.ListUserOrders(userID)
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

func OrderGetProductItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["productID"])

	data := models.ListProductOrders(productID)
	resp := utils.Message(200, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
