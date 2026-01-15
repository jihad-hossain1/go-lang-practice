package handlers

import (
	"ecom/database"
	"ecom/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me an valid id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "P not found", 404)
}
