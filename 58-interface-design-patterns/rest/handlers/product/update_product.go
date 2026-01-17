package product

import (
	"ecom/database"
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "input valid json", 400)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", 201)

}
