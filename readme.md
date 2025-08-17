# Retail Go Backend

A simple backend API built with Go and Gin, providing endpoints for login and reporting attendance, product, and promo data.

## Running Locally

1. **Install dependencies:**
   ```sh
   go mod tidy
   ```

2. **Run the server:**
   ```sh
   go run main.go
   ```

3. The backend will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

### Login

- **POST** `/v1/login`
- **Body:**
  ```json
  {
    "username": "fajar",
    "password": "1234"
  }
  ```

### Report

- **POST** `/v1/report/attendance`
  - **Body:** `{ "status": "hadir" }` or `{ "status": "tidak hadir" }`
- **POST** `/v1/report/product`
  - **Body:** `{ "status": "tersedia" }` or `{ "status": "tidak tersedia" }`
- **POST** `/v1/report/promo`
  - **Body:**
    ```json
    {
      "nama": "Promo Name",
      "harga_normal": 10000,
      "harga_promo": 8000
    }
    ```