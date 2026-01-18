package product

import (
	"ecom/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me an valid id", 400)
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, 404, "P not found")
		return
	}

	util.SendData(w, product, 200)

}
