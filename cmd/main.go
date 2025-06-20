package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"url-shortener/store"
	"url-shortener/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("MONGO_URI not set in .env")
	}

	store.InitMongoDB(mongoURI)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Root)
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/redirect/", handler.Redirect)
	mux.HandleFunc("/analytics/", handler.Analytics)

	fmt.Println("Server running on port 3000...")
	// Wrap the mux with CORS
	if err := http.ListenAndServe(":3000", handler.EnableCORS(mux)); err != nil {
		fmt.Println("Server error:", err)
	}
}
