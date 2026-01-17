package product

import (
	"ecom/database"
	"ecom/util"
	"net/http"
	"strconv"
	// "strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	database.Delete(pId)

	util.SendData(w, "Successfully delete product", 201)

}
