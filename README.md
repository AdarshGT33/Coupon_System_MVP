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
