package handlers

import (
	"ecom/database"
	"ecom/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// parse jwt
	// parse header and payload or claims
	// hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
	// parse signature part from the jwt
	// if the signature and hash is same => forward to create products
	// otherwise 401 status code with Unauthorized

	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := headerArr[1]

	fmt.Println("---token---", accessToken)

	tokenParts := strings.Split(accessToken, ".")

	fmt.Println(tokenParts)

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "input valid json", 400)
		return
	}

	createProduct := database.Store(newProduct)

	util.SendData(w, createProduct, 201)
}
