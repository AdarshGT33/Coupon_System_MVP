
# 🧾 Coupon System API

A lightweight backend service built in **Go (Golang)** to manage and validate coupon codes. The service is containerized using **Docker** for seamless deployment and portability.

---

## 🚀 Tech Stack

- **Backend:** Go (Gin web framework)
- **Database:** PostgreSQL
- **ORM:** GORM
- **Cache:** In-memory (Map)
- **Docs:** Swagger UI
- **Containerization:** Docker (multi-stage build)

---

## 📦 Docker Setup

### 🛠 Prerequisites

- Docker installed
- A running PostgreSQL instance (can be a Docker container or native install)

---

### 🧪 Step-by-step to Run the App

1. **Create `.env` in your project root** (don’t commit it):

```env
DB_URI=host=localhost port=5432 dbname=coupon_db user=postgres password=yourpassword sslmode=disable




## Build the Docker image

```bash
  docker build -t coupon_system .
```


## Run The Container
Make sure your database is running and reachable then

```bash
docker run -p 8080:8080 -v ${PWD}/.env:/root/.env coupon_system
```

If you are on ###Windows- Powershell, use
```powershell
docker run -p 8080:8080 -v ${PWD}/.env:/root/.env coupon_system
```

## API Endpoints

#### Get (health check)

```http
  GET 
```

#### Response
```json
{
    "message":"we are up"
}
```

```http
  POST /create-coupon
```
Create a new coupon

#### Response 
```json
{
  "message": "Coupon created successfully",
  "coupon": {
    "code": "FIRST50",
    "discount": 50,
    "expiresAt": "2025-12-31T23:59:59Z"
  }
}
```

```http
POST /validate-coupon
```
Validate the user's coupon

#### Request
```json
{
    "user_id":"UserID",
    "coupon_code": "DISCOUNT30",   
    "medicine_ids":[
        "medicineid1"
    ],
    "order_value": 500.0,
    "order_time":"2025-05-19T14:30:00Z"
}
```

#### Response

```json
{
		"message":        "Coupon applied successfully",
		"coupon_code":    coupon.CouponCode,
		"discount":       discount,
		"final_amount":   finalAmount,
		"original_price": user_coupon.OrderValue,
}
```



## 📚 Swagger Docs

#### URL:
```http
GET /swagger/index.html
```

Interactive UI to test endpoints and view schema.

## Folder Structure

```go
.
├── main.go
├── controller/
├── utils/
├── cache/
├── docs/
├── go.mod / go.sum
├── Dockerfile
└── .env (not committed)
```
## 🧠 Notes

1. The .env is mounted at runtime and not copied into the image (for security).

2. If your DB is running in Docker, use host.docker.internal instead of localhost on Mac/Windows.

3. Uses a multi-stage Docker build for a small, secure runtime image (~67MB).


## Author

### Adarsh Singh Tomar
Backend Developer | Grit-First Engineer\
Made with ❤️, 🔥, and a Go compiler.




