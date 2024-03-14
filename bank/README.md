# Bank API

This is a simple bank API that simulates payment transaction between two accounts.

## Endpoints

### 1. Transfer
- **Endpoint:** POST `/payment`
- **Description:** Processes a bank transfer.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Request Body:**
  ```json
  {
    "origin_bank": "bbva",
    "origin_account": "123456",
    "destination_bank": "bbva",
    "destination_account": "654321",
    "amount": 178
  }
  ```

### 2. Get Account
- **Endpoint:** POST `/get-account`
- **Description:** Retrieves a customer's account details.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Request Body:**
  ```json
  {
    "bank_name": "bbva",
    "number": "123456",
    "owner_name": "John Nommensen"
  }
  ```