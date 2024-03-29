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

func CheckKey(enteredKey string, client *mongo.Client) bool {
	keyCollection := client.Database("PasswordManager").Collection("Key")

	// Assuming the key is stored as a plain string in the "password" field
	filter := bson.M{"password": DerivePassword(enteredKey)}

	// Use CountDocuments to check if a document with the entered key exists
	count, err := keyCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error checking key in MongoDB:", err)
		return false
	}
	fmt.Println("Count:", count)
	// If count is greater than 0, the key exists
	return count > 0
}

func RetrieveByPlatform(platform string, client *mongo.Client) []string {
	collection := client.Database("PasswordManager").Collection("Passwords")

	filter := bson.M{"platform": platform}

	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error retrieving document:", err)
		return nil
	}

	return []string{result.ID, result.Platform, result.Username, Decrypt(result.Password)}
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

func PrintAllRecords(client *mongo.Client) [][]string {
	collection := client.Database("PasswordManager").Collection("Passwords")

	// Retrieve all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error retrieving documents:", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	// Initialize a slice to store the results
	var results [][]string

	// Iterate through the documents and append values to the results slice
	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Error decoding document:", err)
			continue
		}

		// Append values to the results slice
		result := []string{
			user.Platform,
			user.Username,
			Decrypt(user.Password),
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
		return nil
	}

	// Return the results slice
	return results
}
