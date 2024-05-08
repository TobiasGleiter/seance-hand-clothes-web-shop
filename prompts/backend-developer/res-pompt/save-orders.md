Sure, let's add the functionality to retrieve an order from the client and save it to the database. Here's the updated Go code for the backend:

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Order structure
type Order struct {
    ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name       string             `json:"name,omitempty" bson:"name,omitempty"`
    TotalPrice float64            `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
    UserID     primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
    OrderDate  time.Time          `json:"orderDate,omitempty" bson:"orderDate,omitempty"`
    Articles   []OrderItem        `json:"articles,omitempty" bson:"articles,omitempty"`
}

// OrderItem structure
type OrderItem struct {
    ArticleID primitive.ObjectID `json:"articleId,omitempty" bson:"articleId,omitempty"`
    Quantity  int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

// MongoDB connection information
const (
    connectionString = "mongodb://localhost:27017"
    dbName           = "clothes_db"
    collectionName   = "orders"
)

// Function to save order to database
func saveOrder(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var order Order
    json.NewDecoder(r.Body).Decode(&order)

    // Set order date
    order.OrderDate = time.Now()

    // Insert order into database
    collection := client.Database(dbName).Collection(collectionName)
    result, err := collection.InsertOne(context.Background(), order)
    if err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(w).Encode(result.InsertedID)
}

func main() {
    connectDB()
    defer client.Disconnect(context.Background())

    // Define API endpoints
    http.HandleFunc("/orders", saveOrder)

    // Start the server
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

And here's the updated API documentation:

### API Documentation

#### 1. Save order

- **URL:** `/orders`
- **Method:** POST
- **Description:** Saves an order to the database.
- **Request Body:** Order object.
  ```json
  {
    "name": "String",
    "totalPrice": "Number",
    "userId": "String", // ObjectId of the user
    "articles": [
      {
        "articleId": "String", // ObjectId of the article
        "quantity": "Number"
      }
    ]
  }
  ```
- **Response:**
  - **Status:** 200 OK
  - **Body:** ObjectId of the inserted order.

This API documentation should provide the frontend developer with the necessary information to interact with the new endpoint for saving orders to the database. Let me know if you need further assistance!
