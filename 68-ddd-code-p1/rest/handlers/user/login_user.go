package user

import (
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data ", http.StatusBadRequest)
	}

	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if usr == nil {
		util.SendError(w, http.StatusBadRequest, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateJwt(util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		Email:     usr.Email,
	}, h.cnf.JwtSecretKey)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, accessToken, http.StatusAccepted)
}
