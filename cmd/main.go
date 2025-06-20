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
    // Load the .env file
    if err := godotenv.Load(); err != nil {
        fmt.Println("Warning: No .env file found")
    }

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        panic("MONGO_URI not set in .env")
    }

    // Use the mongoURI variable instead of hardcoded string
    store.InitMongoDB(mongoURI)
    
    http.HandleFunc("/", handler.Root)
    http.HandleFunc("/shorten", handler.Shorten)
    http.HandleFunc("/redirect/", handler.Redirect)
    http.HandleFunc("/analytics/", handler.Analytics)

    fmt.Println("Server running on port 3000...")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Println("Server error:", err)
    }
}