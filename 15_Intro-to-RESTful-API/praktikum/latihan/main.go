package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID int
	Title string
	Content string
}

var data = []article {
	{1, "lorem", "lorem"},
	{2, "ipsum", "ipsum"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", articles)

	// setup server
	var address = "localhost:8000"

	fmt.Printf("server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}