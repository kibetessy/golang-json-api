# Go API Learning Project 🚀

A beginner-friendly HTTP server built with Go (Golang) to learn the fundamentals of building RESTful APIs, working with structs, and returning JSON responses.

## 📚 What I Learned

This project helped me understand:

- Setting up a basic HTTP server in Go
- Using structs to represent and organize data
- Converting Go structs to JSON responses
- Handling multiple API endpoints
- Setting HTTP headers properly
- Error handling in Go
- Testing APIs in the browser

## 🛠️ Technologies Used

- **Go 1.21+** - Programming language
- **net/http** - Go's built-in HTTP server package
- **encoding/json** - JSON encoding/decoding
- **Claude AI** - Prompting and code generation

## 🚀 Prompts used in development
### Prompt 1
I'm building a Go API and currently have a /products endpoint that returns all products. 
I want to learn how to:

1. Create a route that accepts a product ID in the URL (like /products/1 or /products/2)
2. Extract that ID from the URL path
3. Search through my products slice to find the matching product
4. Return that single product as JSON
5. Handle errors when:
   - The ID is not a valid number
   - The product doesn't exist (return 404)

My current code structure:
- I have a Product struct with ID, Name, Price, and InStock fields
- I have a products slice with sample data
- I'm using the standard net/http package (no external frameworks)

Please explain:
- How URL routing works in Go without a router library
- The difference between /products and /products/ (trailing slash)
- How to extract and validate the ID from the URL
- Best practices for error responses in JSON format
- Common mistakes beginners make with URL parsing

Show me the code step-by-step with clear explanations for each part, 
and include examples of testing it in the browser and with curl.

### Prompt 2
I have a working Go API that returns products as JSON with GET requests. 
Now I want to learn how to accept POST requests to add new products to my API.

What I need to understand:

1. How to check if the incoming request is POST vs GET
2. How to read JSON data from the request body
3. How to convert (decode) that JSON into my Product struct
4. How to validate the incoming data:
   - Check if required fields are present
   - Ensure price is positive
   - Make sure name isn't empty
5. How to add the new product to my products slice
6. How to return the created product with a 201 status code
7. What happens to the data when I restart the server (since it's in-memory)

My current setup:
- Using standard library (net/http, encoding/json)
- Products stored in a slice (var products = []Product{...})
- Already have GET /products working
- Want to use the SAME /products endpoint for both GET and POST

Please explain:
- The difference between json.Encoder and json.Decoder
- Why we use r.Body and what it represents
- How to structure the handler to handle multiple HTTP methods
- How to test POST requests (I can't use the browser easily for this)
- Common validation patterns in Go
- Best practices for API responses when creating resources

Include curl examples for testing, explain what each flag means,
and show me what success and error responses should look like.

## 📂 Project Structure

```
go-api-project/
├── main.go           # Main application file with HTTP server and handlers
├── go.mod            # Go module file
└── README.md         # Project documentation
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher installed on your machine
- A web browser or API testing tool (Postman, curl, etc.)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-api-project.git
cd go-api-project
```

2. Initialize the Go module (if not already done):
```bash
go mod init go-api-project
```

3. Run the server:
```bash
go run main.go
```

4. You should see:
```
🚀 Server starting on http://localhost:8080
📍 Endpoints:
   - http://localhost:8080/
   - http://localhost:8080/products
   - http://localhost:8080/health
```

## 🌐 API Endpoints

### 1. Home - `GET /`
Returns a welcome message with available endpoints.

**Example:**
```bash
curl http://localhost:8080/
```

**Response:**
```
Welcome to the Products API!

Available endpoints:
- /products
- /health
```

---

### 2. Get All Products - `GET /products`
Returns a JSON array of all products.

**Example:**
```bash
curl http://localhost:8080/products
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "price": 999.99,
    "in_stock": true
  },
  {
    "id": 2,
    "name": "Mouse",
    "price": 29.99,
    "in_stock": true
  },
  {
    "id": 3,
    "name": "Keyboard",
    "price": 79.99,
    "in_stock": false
  }
]
```

---

### 3. Health Check - `GET /health`
Checks if the API is running properly.

**Example:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "message": "API is running"
}
```

## 💻 Code Overview

### Product Struct

The core data structure representing a product:

```go
type Product struct {
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    InStock bool    `json:"in_stock"`
}
```

### Key Concepts Demonstrated

1. **HTTP Server Setup**: Using `http.ListenAndServe()` to start the server
2. **Route Handling**: Using `http.HandleFunc()` to register endpoints
3. **JSON Encoding**: Using `json.NewEncoder().Encode()` to send JSON responses
4. **HTTP Headers**: Setting `Content-Type: application/json`
5. **Error Handling**: Checking for errors and returning appropriate status codes

## 🧪 Testing the API

### Using a Web Browser
Simply visit the endpoints in your browser:
- http://localhost:8080/
- http://localhost:8080/products
- http://localhost:8080/health

### Using curl
```bash
# Get all products
curl http://localhost:8080/products

# Pretty print JSON
curl http://localhost:8080/products | python -m json.tool
```

### Using Postman or Insomnia
1. Create a new GET request
2. Enter the URL: `http://localhost:8080/products`
3. Send the request
4. View the JSON response

## 📝 What's Next?

Future improvements I'm planning to add:

- [ ] Get single product by ID (`GET /products/:id`)
- [ ] Add new products (`POST /products`)
- [ ] Update existing products (`PUT /products/:id`)
- [ ] Delete products (`DELETE /products/:id`)
- [ ] Connect to a real database (PostgreSQL/MongoDB)
- [ ] Add input validation
- [ ] Implement authentication
- [ ] Add unit tests
- [ ] Deploy to a cloud platform

## 🤔 Challenges I Faced

1. **Understanding struct tags**: Learned why `json:"field_name"` tags are needed for proper JSON formatting
2. **Error handling**: Discovered Go's explicit error handling pattern with multiple return values
3. **Headers**: Realized the importance of setting `Content-Type` header for proper JSON responses

## 📖 Resources

Resources that helped me learn:

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go HTTP Server Tutorial](https://go.dev/doc/tutorial/web-service-gin)
- [JSON and Go](https://go.dev/blog/json)

## 🤝 Contributing

This is a learning project, but suggestions and improvements are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Commit your changes (`git commit -am 'Add some improvement'`)
4. Push to the branch (`git push origin feature/improvement`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Your Name**
- GitHub: [@yourusername](https://github.com/yourusername)
- Learning journey: Building APIs with Go

---

⭐ If you found this project helpful for learning Go, please give it a star!

## 🎯 Learning Goals Achieved

- [x] Set up a Go development environment
- [x] Create a basic HTTP server
- [x] Define and use structs
- [x] Work with JSON encoding
- [x] Handle multiple routes
- [x] Implement error handling
- [x] Test API endpoints

**Status**: ✅ Basic API Complete | 🚧 Advanced Features In Progress
