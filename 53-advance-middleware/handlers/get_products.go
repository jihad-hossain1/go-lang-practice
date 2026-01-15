package handlers

import (
	"ecom/database"
	"ecom/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
