package user

import (
	"ecom/config"
	"ecom/database"
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
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data ", http.StatusBadRequest)
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credential", http.StatusBadRequest)
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		Email:     usr.Email,
	}, cnf.JwtSecretKey)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, accessToken, http.StatusAccepted)
}
