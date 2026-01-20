package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

const (
	totalRequests = 100000
	workers       = 50
)

func worker(wg *sync.WaitGroup, jobs <-chan int, client *http.Client, token string) {
	defer wg.Done()

	for i := range jobs {
		product := Product{
			Title:       fmt.Sprintf("Product %d", i),
			Description: fmt.Sprintf("Description %d", i),
			Price:       float64(i % 100),
			ImageURL:    "https://example.com/image.png",
		}

		body, _ := json.Marshal(product)

		req, _ := http.NewRequest(
			"POST",
			"http://localhost:3010/products",
			bytes.NewBuffer(body),
		)

		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Request error:", err)
			continue
		}
		resp.Body.Close()

		if i%5000 == 0 {
			fmt.Println("Inserted:", i)
		}
	}
}

func main() {
	token := "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImZpcnN0TmFtZSI6ImFiYyIsImVtYWlsIjoiYWJjQGdtYWlsLmNvbSJ9.zKbrw_7c3liaXSdvFtsZROMx8YdjjF_FROsBoITevcI"

	client := &http.Client{}
	jobs := make(chan int, totalRequests)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go worker(&wg, jobs, client, token)
	}

	// Send jobs
	for i := 1; i <= totalRequests; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	fmt.Println("✅ Finished sending 100,000 products")
}

