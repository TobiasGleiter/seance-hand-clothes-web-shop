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
	SellerID   primitive.ObjectID `json:"sellerId,omitempty" bson:"sellerId,omitempty"`
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

// User structure
type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	HashedPassword string            `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
	Salt          string             `json:"salt,omitempty" bson:"salt,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Address       string             `json:"address,omitempty" bson:"address,omitempty"`
	Role          string             `json:"role,omitempty" bson:"role,omitempty"`
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

// Function to read article details from the database
func getArticleDetails(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Assuming article ID is passed in the request URL query parameters
    articleID := r.URL.Query().Get("articleId")
    if articleID == "" {
        http.Error(w, "Missing article ID", http.StatusBadRequest)
        return
    }

    // Convert article ID to ObjectID
    objID, err := primitive.ObjectIDFromHex(articleID)
    if err != nil {
        http.Error(w, "Invalid article ID", http.StatusBadRequest)
        return
    }

    // Find the article in the database
    var article Article
    collection := client.Database(dbName).Collection(collectionName)
    err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&article)
    if err != nil {
        http.Error(w, "Article not found", http.StatusNotFound)
        return
    }

    // Encode the article details and send as response
    json.NewEncoder(w).Encode(article)
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

// Function to register a new user
func registerUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Check if the user already exists
    existingUser := User{}
    collection := client.Database(dbName).Collection(collectionName)
    err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        // User already exists
        http.Error(w, "User already exists", http.StatusBadRequest)
        return
    }

    // Save the new user to the database
    _, err = collection.InsertOne(context.Background(), user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "User registered successfully")
}

// Function to login a user
func loginUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Check if the user exists
    existingUser := User{}
    collection := client.Database(dbName).Collection(collectionName)
    err := collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
    if err != nil {
        // User not found
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    // Validate password
    // Here you would typically compare the hashed password with the password provided by the user,
    // using a secure password hashing algorithm like bcrypt. For simplicity, I'm omitting that part here.

    fmt.Fprintf(w, "User logged in successfully")
}

// Function to read articles for a given user
func getArticlesForUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Assuming user ID is passed in the request body
    var requestData struct {
        UserID string `json:"userId"`
    }
    json.NewDecoder(r.Body).Decode(&requestData)

    userID, err := primitive.ObjectIDFromHex(requestData.UserID)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var articles []Article
    collection := client.Database(dbName).Collection(collectionName)
    cursor, err := collection.Find(context.Background(), bson.M{"sellerId": userID})
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

// Function to create an article for a given user and save it to the database
func createArticleForUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Assuming article details and user ID are passed in the request body
    var requestData struct {
        UserID string `json:"userId"`
        Article Article `json:"article"`
    }
    json.NewDecoder(r.Body).Decode(&requestData)

    // Set the seller ID for the article
    sellerID, err := primitive.ObjectIDFromHex(requestData.UserID)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    requestData.Article.SellerID = sellerID

    // Save the article to the database
    collection := client.Database(dbName).Collection(collectionName)
    _, err = collection.InsertOne(context.Background(), requestData.Article)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "Article created successfully")
}

// Function to delete an article for a given user from the database
func deleteArticleForUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Assuming article ID and user ID are passed in the request body
    var requestData struct {
        UserID   string `json:"userId"`
        ArticleID string `json:"articleId"`
    }
    json.NewDecoder(r.Body).Decode(&requestData)

    // Convert user ID and article ID to ObjectIDs
    userID, err := primitive.ObjectIDFromHex(requestData.UserID)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    articleID, err := primitive.ObjectIDFromHex(requestData.ArticleID)
    if err != nil {
        http.Error(w, "Invalid article ID", http.StatusBadRequest)
        return
    }

    // Delete the article from the database
    collection := client.Database(dbName).Collection(collectionName)
    _, err = collection.DeleteOne(context.Background(), bson.M{"_id": articleID, "sellerId": userID})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(w, "Article deleted successfully")
}

func main() {
    connectDB()
    defer client.Disconnect(context.Background())

    // Define API endpoints
    http.HandleFunc("/articles", getAllArticles)
    http.HandleFunc("/articles/women", getArticlesByCategory)
    http.HandleFunc("/articles/men", getArticlesByCategory)
    http.HandleFunc("/articles/kids", getArticlesByCategory)
	http.HandleFunc("/articleDetails", getArticleDetails)

	http.HandleFunc("/orders", saveOrder)

	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", loginUser)

    http.HandleFunc("/articlesForUser", getArticlesForUser)
	http.HandleFunc("/createArticleForUser", createArticleForUser)
	http.HandleFunc("/deleteArticleForUser", deleteArticleForUser)

    // Start the server
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
