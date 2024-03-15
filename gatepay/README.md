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
- **Response:** With the information of the payment.
  ```json
  {
    "ID": 2,
    "CreatedAt": "2024-03-14T01:54:24.850657-07:00",
    "UpdatedAt": "2024-03-14T01:54:24.850657-07:00",
    "DeletedAt": null,
    "origin_payment_method_id": 1,
    "destination_payment_method_id": 2,
    "amount": 1000,
    "origin_payment_method": {
        "ID": 1,
        "CreatedAt": "2024-03-14T01:23:31.426332-07:00",
        "UpdatedAt": "2024-03-14T01:23:31.426332-07:00",
        "DeletedAt": null,
        "name": "gold",
        "bank_name": "bbva",
        "account_number": "123456",
        "owner_name": "John Nommensen",
        "owner_email": "jnommensen@gmail.com",
        "customer_id": 2,
        "customer": {
            "ID": 2,
            "CreatedAt": "2024-03-14T01:23:22.398943-07:00",
            "UpdatedAt": "2024-03-14T01:23:22.398943-07:00",
            "DeletedAt": null,
            "name": "John Nommensen",
            "email": "jnommensen@gmail.com",
            "payment_methods": null
        }
    },
    "destination_payment_method": {
        "ID": 2,
        "CreatedAt": "2024-03-14T01:23:36.705745-07:00",
        "UpdatedAt": "2024-03-14T01:23:36.705745-07:00",
        "DeletedAt": null,
        "name": "platinum",
        "bank_name": "bbva",
        "account_number": "222222",
        "owner_name": "Deuna",
        "owner_email": "bbutton@deuna.com",
        "merchant_id": 2,
        "merchant": {
            "ID": 2,
            "CreatedAt": "2024-03-14T01:23:22.406869-07:00",
            "UpdatedAt": "2024-03-14T01:23:22.406869-07:00",
            "DeletedAt": null,
            "name": "Deuna",
            "users": null,
            "items": null
        }
    },
    "items": [
        {
            "ID": 2,
            "CreatedAt": "2024-03-14T01:54:24.851823-07:00",
            "UpdatedAt": "2024-03-14T01:54:24.851823-07:00",
            "DeletedAt": null,
            "payment_id": 2,
            "item_id": 1,
            "quantity": 1,
            "price": 1000,
            "payment": null,
            "item": null
        }
    ],
    "status": "payment_status_approved"
  }
  ```

### 2. Get Payments
- **Endpoint:** GET `/get-payments`
- **Description:** Retrieves processed payments for the authenticated user.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Response:** A list of payments sent (customer) or received (merchant) by the authenticated user.

### 3. Add Payment Method
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
- **Response:** With the information of the payment method.
  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-03-14T01:23:31.426332-07:00",
    "UpdatedAt": "2024-03-14T01:23:31.426332-07:00",
    "DeletedAt": null,
    "name": "gold",
    "bank_name": "bbva",
    "account_number": "123456",
    "owner_name": "John Nommensen",
    "owner_email": "jnommensen@gmail.com",
    "customer_id": 2,
    "customer": {
        "ID": 2,
        "CreatedAt": "2024-03-14T01:23:22.398943-07:00",
        "UpdatedAt": "2024-03-14T01:23:22.398943-07:00",
        "DeletedAt": null,
        "name": "John Nommensen",
        "email": "jnommensen@gmail.com",
        "payment_methods": null
    }
  }
  ```
  
### 4. Get Activity Log
- **Endpoint:** POST `/get-activity-log`
- **Description:** Retrieves the activity log with the paginated, sorting and filters specified.
- **Headers:** Authorization: Bearer {{auth_token}}
- **Request Body:**
  ```json
  {
    "page": 1,
    "page_size": 10,
    "order_by": "created_at desc",
    "filters": [
        {
            "property": "author",
            "value": "jnommensen@gmail.com",
            "relational_operator": "=",
            "logical_operator": "and"
        }
    ]
  }
  ```
- **Response:** A list of activity logs with the specified filters.
  ```json
  [
    {
        "ID": 1,
        "CreatedAt": "2024-03-14T21:19:57.592125-07:00",
        "UpdatedAt": "2024-03-14T21:19:57.592125-07:00",
        "DeletedAt": null,
        "type": "payment_method",
        "author": "jnommensen@gmail.com",
        "action": "create",
        "detail": {
            "account": "bbva-123456",
            "id": 1,
            "name": "gold"
        }
    }
  ]
  ```