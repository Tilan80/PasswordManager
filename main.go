package main

import (
	"context"
	"fmt"
	"math/rand"
)

const uri = "mongodb://localhost:27017"
const dbName = "PasswordManager"
const collectionName = "Passwords"
const keydb = "Key"

type User struct {
	ID       string `bson:"_id,omitempty"`
	Platform string `bson:"platform"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func main() {
	client, err := ConnectToMongoDB(uri)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var enteredKey string
	fmt.Println("Enter the password:")
	fmt.Scanln(&enteredKey)
	//////////////////////
	if CheckKey(DerivePassword(enteredKey), client, dbName, keydb) {
		fmt.Println("Key is correct")
		for {
			var com string
			fmt.Println("Enter the command:\nadd\naddOwn - to add with own password\nget - to get by platform\ngetAll" +
				"\ndel - to delete\nex - to exit")
			fmt.Scanln(&com)
			//////////////////////////////
			if com == "add" {
				var plat, us, pass string
				fmt.Println("Enter platform:")
				fmt.Scanln(&plat)
				fmt.Println("Enter username:")
				fmt.Scanln(&us)
				pass = Encrypt(GenPassword(20))
				user := User{Platform: plat, Username: us, Password: pass}

				err = InsertDocument(client, dbName, collectionName, user)
				if err != nil {
					fmt.Println("Error inserting document:", err)
				} else {
					fmt.Println("Document inserted successfully")
				}
				///////////////////////
			} else if com == "addOwn" {
				var plat, us, pass string
				fmt.Println("Enter platform:")
				fmt.Scanln(&plat)
				fmt.Println("Enter username:")
				fmt.Scanln(&us)
				fmt.Println("Enter password:")
				fmt.Scanln(&pass)
				user := User{Platform: plat, Username: us, Password: Encrypt(pass)}

				err = InsertDocument(client, dbName, collectionName, user)
				if err != nil {
					fmt.Println("Error inserting document:", err)
				} else {
					fmt.Println("Document inserted successfully")
				}
				///////////////////////////////
			} else if com == "get" {
				var plat string
				fmt.Println("Enter platform to retrieve:")
				fmt.Scanln(&plat)
				RetrieveByPlatform(plat, client, dbName, collectionName)
				//////////////////////////////////
			} else if com == "getAll" {
				PrintAllRecords(client, dbName, collectionName)
			} else if com == "ex" {
				break
				//////////////////////////////
			} else if com == "del" {
				var plat string
				fmt.Println("Enter platform to delete:")
				fmt.Scanln(&plat)
				DeleteByPlatform(plat, client, dbName, collectionName)
				//////////////////////////7
			} else {
				fmt.Println("Invalid operation ", com)
			}
		}
	} else {
		fmt.Println("Invalid key")
	}

}

func GenPassword(len int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&'()*+,-./:;<=>?@[]^_`{|}~"

	password := make([]byte, len)
	charLen := 92
	for i := range password {
		password[i] = characters[rand.Intn(charLen)]
	}
	secure_password := string(password)
	fmt.Printf(" secure pass: " + secure_password)

	return secure_password
}
