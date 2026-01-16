package main

import (
	// "ecom/cmd"
	"ecom/util"
	"fmt"
)

func main() {
	// cmd.Serve()
	jwt, err := util.CreateJwt(util.Payload{
		Sub:       1,
		FirstName: "abc",
		Email:     "abc@gmail.com",
	}, "my-secret")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)
}
