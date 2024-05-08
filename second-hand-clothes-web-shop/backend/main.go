package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
	"time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
)

// Article structure
type Article struct {
    ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name       string             `json:"name,omitempty" bson:"name,omitempty"`
    Category   string             `json:"category,omitempty" bson:"category,omitempty"`
    Price      float64            `json:"price,omitempty" bson:"price,omitempty"`
    Size       string             `json:"size,omitempty" bson:"size,omitempty"`
    Rating     float64            `json:"rating,omitempty" bson:"rating,omitempty"`
    Subcategory string            `json:"subcategory,omitempty" bson:"subcategory,omitempty"`
}

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
    collectionName   = "articles"
)

var client *mongo.Client

// Connect to MongoDB
func connectDB() {
    clientOptions := options.Client().ApplyURI(connectionString)
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")
}

// Endpoint to get all articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var articles []Article
    collection := client.Database(dbName).Collection(collectionName)
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var article Article
        err := cursor.Decode(&article)
        if err != nil {
            log.Fatal(err)
        }
        articles = append(articles, article)
    }
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(w).Encode(articles)
}

// Endpoint to get articles by category
func getArticlesByCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    category := r.URL.Query().Get("category")
    var articles []Article
    collection := client.Database(dbName).Collection(collectionName)
    cursor, err := collection.Find(context.Background(), bson.M{"category": category})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var article Article
        err := cursor.Decode(&article)
        if err != nil {
            log.Fatal(err)
        }
        articles = append(articles, article)
    }
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(w).Encode(articles)
}

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
    http.HandleFunc("/articles", getAllArticles)
    http.HandleFunc("/articles/women", getArticlesByCategory)
    http.HandleFunc("/articles/men", getArticlesByCategory)
    http.HandleFunc("/articles/kids", getArticlesByCategory)

	// Define API endpoints
	http.HandleFunc("/orders", saveOrder)

    // Start the server
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
