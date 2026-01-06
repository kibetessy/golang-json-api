
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Product represents an item in our store
type Product struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    InStock bool    `json:"in_stock"`
}

// Our in-memory database (just a slice for now)
var products = []Product{
    {ID: 1, Name: "Laptop", Price: 999.99, InStock: true},
    {ID: 2, Name: "Mouse", Price: 29.99, InStock: true},
    {ID: 3, Name: "Keyboard", Price: 79.99, InStock: false},
}

func main() {
    // Register routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/products", productsHandler)
    http.HandleFunc("/health", healthHandler)
    
    // Start server
    fmt.Println("🚀 Server starting on http://localhost:8080")
    fmt.Println("📍 Endpoints:")
    fmt.Println("   - http://localhost:8080/")
    fmt.Println("   - http://localhost:8080/products")
    fmt.Println("   - http://localhost:8080/health")
    
    http.ListenAndServe(":8080", nil)
}

// homeHandler handles requests to /
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Products API!\n\nAvailable endpoints:\n- /products\n- /health")
}

// productsHandler returns all products as JSON
func productsHandler(w http.ResponseWriter, r *http.Request) {
    // Set response header to JSON
    w.Header().Set("Content-Type", "application/json")
    
    // Encode products to JSON and send response
    err := json.NewEncoder(w).Encode(products)
    if err != nil {
        // If encoding fails, send error response
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
}

// healthHandler checks if the API is running
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    response := map[string]string{
        "status": "ok",
        "message": "API is running",
    }
    
    json.NewEncoder(w).Encode(response)
}
