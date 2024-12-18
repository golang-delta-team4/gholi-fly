### Data Models and APIs for Bank Microservice
<iframe width="100%" height="500px" allowtransparency="true" allowfullscreen="true" scrolling="no" title="Embedded DB Designer IFrame" frameborder="0" src='https://erd.dbdesigner.net/designer/schema/1734530282-bank?embed=true'></iframe>
---

#### **Factor Model**

| **Field**          | **Description**                                                            |
| ------------------ | -------------------------------------------------------------------------- |
| `ID`               | Unique identifier (UUID) for the factor.                                   |
| `SourceService`    | Name of the microservice generating the factor (e.g., HotelService).       |
| `ExternalID`       | Unique identifier for the factor in the source service.                    |
| `Amount`           | Total amount of the factor.                                                |
| `Currency`         | Currency used for the factor (e.g., IRR).                                  |
| `CustomerWalletID` | Wallet ID of the customer associated with the factor.                      |
| `Purpose`          | Purpose of the factor (e.g., HotelBooking, TransportBooking).              |
| `Details`          | JSON object for flexible fields (e.g., check-in date, origin/destination). |
| `BookingID`        | ID of the booking this factor is linked to (for aggregation).              |
| `Status`           | Status of the factor (`Pending`, `Approved`, `Rejected`).                  |
| `CreatedAt`        | Timestamp when the factor was created.                                     |
| `UpdatedAt`        | Timestamp when the factor was last updated.                                |

---

#### **CreditCard Model**

| **Field**    | **Description**                                            |
| ------------ | ---------------------------------------------------------- |
| `ID`         | Unique identifier (UUID) for the credit card.              |
| `WalletID`   | Foreign key referencing the wallet this card is linked to. |
| `CardNumber` | The 16-digit card number following Iranian bank standards. |
| `ExpiryDate` | Expiration date of the card.                               |
| `CVV`        | Security code of the card.                                 |
| `HolderName` | Name of the cardholder.                                    |
| `CreatedAt`  | Timestamp when the card was linked.                        |
| `UpdatedAt`  | Timestamp when the card details were last updated.         |

---

#### **Transaction Model**

| **Field**   | **Description**                                                      |
| ----------- | -------------------------------------------------------------------- |
| `ID`        | Unique identifier (UUID) for the transaction.                        |
| `WalletID`  | Foreign key referencing the wallet involved in the transaction.      |
| `FactorID`  | Foreign key referencing the factor associated with this transaction. |
| `Amount`    | Amount involved in the transaction.                                  |
| `Type`      | Type of the transaction (`Credit`, `Debit`, `Refund`).               |
| `Status`    | Status of the transaction (`Pending`, `Completed`, `Failed`).        |
| `CreatedAt` | Timestamp when the transaction was created.                          |
| `UpdatedAt` | Timestamp when the transaction was last updated.                     |

---

#### **Wallet Model**

| **Field**   | **Description**                                               |
| ----------- | ------------------------------------------------------------- |
| `ID`        | Unique identifier (UUID) for the wallet.                      |
| `OwnerID`   | Reference to the owner of the wallet (user, company, or app). |
| `Type`      | Type of the wallet (`Person`, `Company`, `App`).              |
| `Balance`   | Current balance of the wallet.                                |
| `Currency`  | Currency of the wallet (e.g., IRR).                           |
| `CreatedAt` | Timestamp when the wallet was created.                        |
| `UpdatedAt` | Timestamp when the wallet was last updated.                   |

---

### JSON Requests and Responses

#### **Factor API**

##### **Request: Create Factor**

```json
{
  "Action": "CreateFactor",
  "Factor": {
    "SourceService": "HotelService",
    "ExternalID": "hotel_reservation_123",
    "Amount": 800000,
    "Currency": "IRR",
    "CustomerWalletID": "customer_wallet_001",
    "Purpose": "HotelBooking",
    "Details": {
      "CheckIn": "2024-12-20",
      "CheckOut": "2024-12-25",
      "HotelName": "Luxury Inn"
    },
    "BookingID": "booking_123"
  }
}
```

##### **Response: Create Factor**

- **Success**:

```json
{
  "FactorID": "factor_hotel_001",
  "BookingID": "booking_123",
  "Status": "Approved",
  "ReservedAmount": 800000
}
```

- **Failure**:

```json
{
  "FactorID": "factor_hotel_001",
  "BookingID": "booking_123",
  "Status": "Rejected",
  "Reason": "Insufficient Funds"
}
```

#### **Transaction API**

##### **Request: Distribute Revenue**

```json
{
  "Action": "DistributeRevenue",
  "FactorID": "factor_hotel_001",
  "BookingID": "booking_123",
  "Distribution": [
    { "WalletID": "hotel_company_wallet", "Amount": 600000 },
    { "WalletID": "hotel_employee_1_wallet", "Amount": 200000 }
  ]
}
```

##### **Response: Distribute Revenue**

- **Success**:

```json
{
  "FactorID": "factor_hotel_001",
  "BookingID": "booking_123",
  "Status": "Completed",
  "Details": [
    { "WalletID": "hotel_company_wallet", "Amount": 600000, "Status": "Success" },
    { "WalletID": "hotel_employee_1_wallet", "Amount": 200000, "Status": "Success" }
  ]
}
```

- **Failure**:

```json
{
  "FactorID": "factor_hotel_001",
  "BookingID": "booking_123",
  "Status": "Failed",
  "Details": [
    { "WalletID": "hotel_company_wallet", "Amount": 600000, "Status": "Success" },
    { "WalletID": "hotel_employee_1_wallet", "Amount": 200000, "Status": "Failed", "Reason": "Invalid Wallet" }
  ]
}
```

#### **Wallet API**

##### **Request: Wallet Balance Check**

```json
{
  "Action": "CheckBalance",
  "WalletID": "customer_wallet_001"
}
```

##### **Response: Wallet Balance Check**

- **Success**:

```json
{
  "WalletID": "customer_wallet_001",
  "Balance": 2000000,
  "Currency": "IRR"
}
```

---

#### **Aggregate Factors API**

##### **Request: Aggregate Factors**

```json
{
  "Action": "AggregateFactors",
  "BookingID": "booking_123",
  "Factors": [
    {
      "FactorID": "factor_hotel_001",
      "Amount": 800000,
      "SourceService": "HotelService"
    },
    {
      "FactorID": "factor_transport_002",
      "Amount": 1200000,
      "SourceService": "TransportService"
    }
  ],
  "TotalAmount": 2000000,
  "CustomerWalletID": "customer_wallet_001"
}
```

##### **Response: Aggregate Factors**

- **Success**:

```json
{
  "BookingID": "booking_123",
  "Status": "Approved",
  "ReservedAmount": 2000000
}
```

- **Failure**:

```json
{
  "BookingID": "booking_123",
  "Status": "Rejected",
  "Reason": "Insufficient Funds"
}
```

