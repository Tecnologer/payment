# Auth

It's a pretty simple microservice that generates JWT tokens and validates them.

## Endpoints

### 1. Login
- **Endpoint:** POST `/login`
- **Description:** Authenticates a customer and returns a token.
- **Request Body:**
  ```json
  {
    "username": "<email>",
    "password": "<fictional_password>"
  }
  ```
- **Response:**
  ```json
    {
        "token": "<JWT_TOKEN>"
    }
  ```

### 3. Validate Token
- **Endpoint:** GET `/validate`
- **Description:** Validates the provided token. Responses with 200 and information about the token if it's valid.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Response:**
  ```json
    {
      "exp": <int>,
      "username": "<email>"
    }
  ```
