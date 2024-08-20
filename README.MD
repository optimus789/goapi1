# GO API Service

## Overview

This GO API Service provides endpoints to manage and retrieve coin balances for users. It is built using the Go programming language and leverages the `chi` router for handling HTTP requests.

## Endpoints

### Get Coin Balance

**Endpoint:** `/get-coin-balance`

**Method:** `GET`

**Parameters:**
- `username` (query parameter): The username for which to retrieve the coin balance.

**Response:**
- [`200 OK`](command:_github.copilot.openSymbolFromReferences?%5B%22200%20OK%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22g%3A%5C%5CgoTutorials%5C%5Cgo.sum%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fg%253A%2FgoTutorials%2Fgo.sum%22%2C%22path%22%3A%22%2Fg%3A%2FgoTutorials%2Fgo.sum%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A3%2C%22character%22%3A55%7D%7D%5D%5D "Go to definition"): Returns the coin balance for the specified user.
- [`400 Bad Request`](command:_github.copilot.openSymbolFromReferences?%5B%22400%20Bad%20Request%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22g%3A%5C%5CgoTutorials%5C%5Capi%5C%5Capi.go%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fg%253A%2FgoTutorials%2Fapi%2Fapi.go%22%2C%22path%22%3A%22%2Fg%3A%2FgoTutorials%2Fapi%2Fapi.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A33%2C%22character%22%3A1%7D%7D%5D%5D "Go to definition"): If the request is malformed or missing required parameters.
- [`500 Internal Server Error`](command:_github.copilot.openSymbolFromReferences?%5B%22500%20Internal%20Server%20Error%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22g%3A%5C%5CgoTutorials%5C%5Capi%5C%5Capi.go%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fg%253A%2FgoTutorials%2Fapi%2Fapi.go%22%2C%22path%22%3A%22%2Fg%3A%2FgoTutorials%2Fapi%2Fapi.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A18%2C%22character%22%3A3%7D%7D%2C%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22g%3A%5C%5CgoTutorials%5C%5Ccmd%5C%5Capi%5C%5Cmain.go%22%2C%22_sep%22%3A1%2C%22external%22%3A%22file%3A%2F%2F%2Fg%253A%2FgoTutorials%2Fcmd%2Fapi%2Fmain.go%22%2C%22path%22%3A%22%2FG%3A%2FgoTutorials%2Fcmd%2Fapi%2Fmain.go%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A19%2C%22character%22%3A6%7D%7D%5D%5D "Go to definition"): If an unexpected error occurs.

**Example Request:**
```sh
curl -X GET "http://localhost:8000/get-coin-balance?username=johndoe"
```

**Example Response:**
```json
{
    "code": 200,
    "balance": 1000
}
```

**Error Handling**

The API provides structured error responses with appropriate HTTP status codes.

**Error Response Structure:**
```json
{
    "code": <HTTP_STATUS_CODE>,
    "message": "<ERROR_MESSAGE>"
}
```

**Running the Service:**
To run the GO API Service, execute the following command:
```sh
go run cmd/api/main.go
```
The service will start and listen on 0.0.0.0:8000


**Setting up:**
- Clone the repository.
- Navigate to the project directory.
- Install dependencies using go mod tidy.


**License:**
This project is licensed under the MIT License.