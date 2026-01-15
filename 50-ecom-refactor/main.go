package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Server is ok")
}

/*
- when struct define key First character small letter then not accessble other package ( this are the private method to allow intier main package only)
- we define this way `json:"id"`
*/
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != "GET" { // r.Method = post, put , patch, delete
		http.Error(w, "Invalid Api Method", 400)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	handlePreflightReq(w, r)

	if r.Method != "POST" { // r.Method = post, put , patch, delete
		http.Error(w, "Invalid Api Method", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "input valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

/* Cors issue */
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/", handler) // route

	mux.HandleFunc("/products", getProducts) // route
	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Mange",
		Description: "I Love Mango",
		Price:       110,
		ImgUrl:      "https://c8.alamy.com/comp/2N87A9W/mongo-fruit-hanging-on-tree-mongo-fruits-2N87A9W.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Mango juice",
		Description: "Fresh Mango Juice",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsx6-7fkbFv4qP3a9rsGy3zJSQIbeDamE2Pg&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Fresh organic source",
		Price:       10,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/organic%20bananas.png?itok=_JpbRjWp-xPBdBLll",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)

}
