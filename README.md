# Go API Starter Guide

This guide provides instructions to set up and run a basic Go API. 

## Prerequisites

- Go installed on your machine (version 1.16 or higher recommended)

## Getting Started

1. **Clone the repository**

   ```sh
   git clone https://github.com/joaorfp/GO-api.git
   cd GO-api/

2. **Initialize the Go module**

```bash
 $ go mod tidy
 ```    

3. **Run the API**
```bash
 $ go run cmd/main.go
 ```

### Sample API Requests

1. **Create a User**
```bash
 $ curl -X POST -H "Content-Type: application/json" -d '{"username": "johndoe", "email": "johndoe@example.com"}' http://localhost:8080/users
 ```

2. **Create a Bank Accont**
```bash
 $ curl -X POST -H "Content-Type: application/json" -d '{"balance": 1000.50, "invested_balance": 500.00, "user_id": 1}' http://localhost:8080/accounts
 ```

3. **Get All Users**
```bash
 $ curl -X GET http://localhost:8080/users
 ```

4. **Get All Bank Accounts**
```bash
 $ curl -X GET http://localhost:8080/accounts
 ```