package user

import (
	"ecom/repo"
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req RequestUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data ", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, usr, http.StatusCreated)
}
