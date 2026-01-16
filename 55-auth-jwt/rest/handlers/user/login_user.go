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

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credential", http.StatusBadRequest)
	}

	util.SendData(w, usr, http.StatusAccepted)
}
