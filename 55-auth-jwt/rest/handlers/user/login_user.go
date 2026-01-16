package user

import (
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data ", http.StatusBadRequest)
	}

	createUser := newUser.Store()
	util.SendData(w, createUser, http.StatusCreated)
}
