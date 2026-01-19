package product

import (
	"ecom/domain"
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req RequestProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "input valid json", 400)
		return
	}

	createProduct, err := h.svc.Create(domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	}) //database.Store(newProduct)

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	util.SendData(w, createProduct, 201)
}
