# Designed data structure by ChatGPT

## Articles:

```json
{
  "_id": ObjectId,
  "name": "String",
  "category": "String", // "men", "women", "kids"
  "price": "Number",
  "size": "String",
  "rating": "Number",
  "subcategory": "String",
  "sellerId": ObjectId, // "new, forgot this in the prompt
}
```

## Order:

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

## User:

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
