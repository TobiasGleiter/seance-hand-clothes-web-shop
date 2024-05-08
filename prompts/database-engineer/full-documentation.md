# Full ChatGPT generated documentation for the database.

### MongoDB Data Structure Documentation

#### Articles Collection

```json
{
  "_id": ObjectId,
  "name": "String",
  "category": "String", // "men", "women", "kids"
  "price": "Number",
  "size": "String",
  "rating": "Number",
  "subcategory": "String"
}
```

- **\_id**: Unique identifier for the article.
- **name**: Name of the article.
- **category**: Category of the article ("men", "women", "kids").
- **price**: Price of the article.
- **size**: Size of the article (e.g., "S", "M", "L").
- **rating**: Rating of the article.
- **subcategory**: Subcategory of the article (optional).

#### Orders Collection

```json
{
  "_id": ObjectId,
  "name": "String",
  "totalPrice": "Number",
  "userId": ObjectId,
  "orderDate": Date,
  "articles": [
    {
      "articleId": ObjectId,
      "quantity": Number
    }
  ]
}
```

- **\_id**: Unique identifier for the order.
- **name**: Name of the order (optional).
- **totalPrice**: Total price of the order.
- **userId**: Reference to the user who placed the order.
- **orderDate**: Date when the order was placed.
- **articles**: Array of objects containing article information and quantity ordered.
  - **articleId**: Reference to the article.
  - **quantity**: Quantity of the article ordered.

#### Users Collection

```json
{
  "_id": ObjectId,
  "name": "String",
  "hashedPassword": "String",
  "salt": "String",
  "email": "String",
  "address": "String",
  "role": "String" // "user", "admin"
}
```

- **\_id**: Unique identifier for the user.
- **name**: Name of the user.
- **hashedPassword**: Hashed password of the user.
- **salt**: Salt used for hashing the password.
- **email**: Email address of the user.
- **address**: Address of the user.
- **role**: Role of the user ("user", "admin").

### Used Database

For the backend-developer team, you can use MongoDB as the database to implement this data structure. Ensure to securely hash passwords using a strong hashing algorithm and use unique salts for each user to enhance security. Additionally, consider data validation and appropriate indexes for efficient querying. Let me know if you need further assistance or modifications!
