package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"i am abc.")
}

func handler (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Fprintf(w,"Server is ok")
}

func main (){
	mux := http.NewServeMux()

	mux.HandleFunc("/",handler)
	
	mux.HandleFunc("/hello",helloHandler)

	mux.HandleFunc("/about",aboutHandler)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000",mux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

