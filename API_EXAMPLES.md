# API Examples

This document provides example API requests and responses for the Go Auth service.

## Base URL

```
http://localhost:3000
```

## Authentication

All user endpoints require JWT authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### 1. Register User

**POST** `/auth/signup`

```bash
curl -X POST http://localhost:3000/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "phoneNumber": "09123456789"
  }'
```

**Response:**

```json
{
  "succeed": true,
  "message": "OTP sent successfully",
  "otpCode": "123456"
}
```

### 2. Confirm OTP

**POST** `/auth/signup/confirm-otp`

```bash
curl -X POST http://localhost:3000/auth/signup/confirm-otp \
  -H "Content-Type: application/json" \
  -d '{
    "phoneNumber": "09123456789",
    "otpCode": "123456"
  }'
```

**Response:**

```json
{
  "succeed": true,
  "message": "User registered successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Get Current User

**GET** `/user/`

```bash
curl -X GET http://localhost:3000/user/ \
  -H "Authorization: Bearer <your_jwt_token>"
```

**Response:**

```json
{
  "succeed": true,
  "message": "Succeed",
  "user": {
    "id": "689b9bbf5800ec55229e240b",
    "phoneNumber": "09123456789",
    "createdDate": "2025-08-12T19:53:35.685Z"
  }
}
```

### 4. Get User By ID

**GET** `/user/{userId}`

```bash
curl -X GET http://localhost:3000/user/689b9bbf5800ec55229e240b \
  -H "Authorization: Bearer <your_jwt_token>"
```

**Response:**

```json
{
  "succeed": true,
  "message": "Succeed",
  "user": {
    "id": "689b9bbf5800ec55229e240b",
    "phoneNumber": "09123456789",
    "createdDate": "2025-08-12T19:53:35.685Z"
  }
}
```

### 5. Get All Users with Pagination

**GET** `/user/all`

#### Basic Request (Default pagination)

```bash
curl -X GET "http://localhost:3000/user/all" \
  -H "Authorization: Bearer <your_jwt_token>"
```

#### With Pagination Parameters

```bash
curl -X GET "http://localhost:3000/user/all?page=1&pageSize=5" \
  -H "Authorization: Bearer <your_jwt_token>"
```

#### Search by Phone Number

```bash
curl -X GET "http://localhost:3000/user/all?phone=0912" \
  -H "Authorization: Bearer <your_jwt_token>"
```

#### Combined: Pagination + Search

```bash
curl -X GET "http://localhost:3000/user/all?page=2&pageSize=10&phone=091" \
  -H "Authorization: Bearer <your_jwt_token>"
```

**Response:**

```json
{
  "succeed": true,
  "message": "Success",
  "users": [
    {
      "id": "689b9bbf5800ec55229e240b",
      "phoneNumber": "09123456789",
      "createdDate": "2025-08-12T19:53:35.685Z"
    },
    {
      "id": "689b9bbf5800ec55229e240c",
      "phoneNumber": "09187654321",
      "createdDate": "2025-08-12T20:15:22.123Z"
    }
  ],
  "totalCount": 25,
  "page": 1,
  "pageSize": 10,
  "totalPages": 3
}
```

## Query Parameters for Get All Users

| Parameter | Type   | Required | Default | Description                                    |
|-----------|--------|----------|---------|------------------------------------------------|
| page      | int    | No       | 1       | Page number (starts from 1)                   |
| pageSize  | int    | No       | 10      | Number of users per page (max: 100)           |
| phone     | string | No       | -       | Search users by phone number (partial match)  |

## Error Responses

### 400 Bad Request

```json
{
  "succeed": false,
  "message": "Invalid request parameters"
}
```

### 401 Unauthorized

```json
{
  "succeed": false,
  "message": "Invalid or missing JWT token"
}
```

### 404 Not Found

```json
{
  "succeed": false,
  "message": "User Not Found"
}
```

## Examples with JavaScript (Fetch API)

### Get All Users with Pagination

```javascript
const token = 'your_jwt_token_here';

// Basic request
fetch('http://localhost:3000/user/all', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(response => response.json())
.then(data => console.log(data));

// With search and pagination
const params = new URLSearchParams({
  page: '1',
  pageSize: '5',
  phone: '0912'
});

fetch(`http://localhost:3000/user/all?${params}`, {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(response => response.json())
.then(data => {
  console.log('Users:', data.users);
  console.log('Total:', data.totalCount);
  console.log('Current Page:', data.page);
  console.log('Total Pages:', data.totalPages);
});
```

## Rate Limiting

- OTP requests are limited to 3 attempts per phone number within 10 minutes
- If limit is exceeded, you'll receive:

```json
{
  "succeed": false,
  "message": "Rate limit exceeded. Try again later."
}
```

## Notes

1. All timestamps are in ISO 8601 format (UTC)
2. Phone numbers should be in the format: `09XXXXXXXXX` (Iranian format)
3. OTP codes are 4-digit numbers
4. OTP codes expire after 2 minutes
5. JWT tokens should be stored securely and included in all authenticated requests
6. The phone search is case-insensitive and supports partial matching
7. Maximum page size is 100 users per request
