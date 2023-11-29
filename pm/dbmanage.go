package pm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Platform string `bson:"platform"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type App struct{}

func ConnectToMongoDB(ctx context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InsertDocument(client *mongo.Client, user User) error {
	collection := client.Database("PasswordManager").Collection("Passwords")

	opts := options.InsertOne().SetBypassDocumentValidation(true)

	_, err := collection.InsertOne(context.TODO(), user, opts)
	return err
}

func CheckKey(app *App, enteredKey string, client *mongo.Client) bool {
	keyCollection := client.Database("PasswordManager").Collection("Key")

	// Assuming the key is stored as a plain string in the "password" field
	filter := bson.M{"password": DerivePassword(enteredKey)}

	// Use CountDocuments to check if a document with the entered key exists
	count, err := keyCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error checking key in MongoDB:", err)
		return false
	}

	// If count is greater than 0, the key exists
	return count > 0
}

func RetrieveByPlatform(platform string, client *mongo.Client) {
	collection := client.Database("PasswordManager").Collection("Passwords")

	filter := bson.M{"platform": platform}

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error retrieving document:", err)
		return
	}

	fmt.Printf("Platform: %s\nUsername: %s\nPassword: %s\n", result.Platform, result.Username, Decrypt(result.Password))
}

func DeleteByPlatform(platform string, client *mongo.Client) {
	collection := client.Database("PasswordManager").Collection("Passwords")

	filter := bson.M{"platform": platform}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error deleting document:", err)
		return
	}

	if result.DeletedCount == 0 {
		fmt.Println("No document found for deletion.")
		return
	}

	fmt.Printf("Deleted %d document(s) for platform %s\n", result.DeletedCount, platform)
}

func PrintAllRecords(client *mongo.Client) {
	collection := client.Database("PasswordManager").Collection("Passwords")

	// Retrieve all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error retrieving documents:", err)
		return
	}
	defer cursor.Close(context.TODO())

	// Iterate through the documents and print platform, username, and password
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}

		fmt.Printf("Platform: %s, Username: %s, Password: %s\n", user.Platform, user.Username, Decrypt(user.Password))
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return
	}
}
