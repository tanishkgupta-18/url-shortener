package main

import (
	"fmt"
	"net/http"
	"url-shortener/handler"
	"url-shortener/store"
)

func main() {
	store.InitMongoDB("mongodb://localhost:27017")

	http.HandleFunc("/", handler.Root)
	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/redirect/", handler.Redirect)
	http.HandleFunc("/analytics/", handler.Analytics)

	fmt.Println("Server running on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
