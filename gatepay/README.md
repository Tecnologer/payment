# GatePay

Online payment platform API-based application focused on e-commerce businesses to securely and seamlessly process transactions.

## Endpoints

### 1. Pay
- **Endpoint:** POST `/pay`
- **Description:** Initiates a payment process for the authenticated user.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Request Body:**
  ```json
  {
    "origin_payment_method_id": 1,
    "destination_payment_method_id": 2,
    "amount": 1000.00,
    "items": [
      {
        "description": "SmartTV 50 inches",
        "price": 1000.00,
        "quantity": 1
      }
    ]
  }
  ```

### 2. Get Payments
- **Endpoint:** GET `/get-payments`
- **Description:** Retrieves processed payments for the authenticated user.
- **Headers:** Authorization: Bearer {{auth_token}}

### 3. Add Payment Method - Customer
- **Endpoint:** POST `/add-payment-method`
- **Description:** Adds a customer's payment method to the authenticated user.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Request Body:**
  ```json
  {
    "bank_name": "bbva",
    "account_number": "123456",
    "name": "gold"
  }
  ```