# Go Fiber Products with JWT Authentication

## Requirements
- Go
- PostgreSQL
- Air (optional)

## How to run

If you have Air installed, run command below:

```bash
air
```

If you don't have Air installed, run command below:

```bash
go run main.go
```

## .ENV SETUP

Create a .env file in the root directory of the project and add the following environment variables:

```
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=go-fiber-modular
DB_TIMEZONE=Asia/Jakarta
PORT=3000
JWT_SECRET=secret
```

## API ENDPOINTS
🔒 = Requires authentication

### AUTH
| Method   | Endpoint         | Function Name | Description                                  |
| -------- | ---------------- | ------------- | -------------------------------------------- |
| **POST** | `/auth/register` | `Register`    | Register a new user (Name, Email, Password). |
| **POST** | `/auth/login`    | `Login`       | Login and receive JWT token.                 |


### TOKO
| Method   | Endpoint    | Function Name | Description                                         |
| -------- | ----------- | ------------- | --------------------------------------------------- |
| **POST** | `/toko`     | `CreateToko`  | 🔒 Create a new toko. `UserID` is taken from token. |
| **GET**  | `/toko/my`  | `GetMyToko`   | 🔒 Get list of toko owned by logged-in user.        |
| **PUT**  | `/toko/:id` | `UpdateToko`  | 🔒 Update toko data by ID.                          |


### PRODUCTS
| Method   | Endpoint    | Function Name | Description                                         |
| -------- | ----------- | ------------- | --------------------------------------------------- |
| **POST** | `/toko`     | `CreateToko`  | 🔒 Create a new toko. `UserID` is taken from token. |
| **GET**  | `/toko/my`  | `GetMyToko`   | 🔒 Get list of toko owned by logged-in user.        |
| **PUT**  | `/toko/:id` | `UpdateToko`  | 🔒 Update toko data by ID.                          |
