package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Product struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
	Category string `json:"category"`
}

func getRequestData() []byte {
	resp, _ := http.Get("https://fakestoreapi.com/products")
	body, _ := io.ReadAll(resp.Body)

	return body
}

func main() {

	var wg sync.WaitGroup
	var products []Product

	wg.Add(1)
	go func() {
		json.Unmarshal(getRequestData(), &products)

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("products data")
	fmt.Println("===")
	for _, product := range products {
		fmt.Println("title:", product.Title)
		fmt.Println("price:", product.Price)
		fmt.Println("category:", product.Category)
		fmt.Println("===")
	}

	// sb := string(body)
	// fmt.Println(sb)
}