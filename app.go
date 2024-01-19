package main

import (
	"PasswordManager/pm"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient = new(mongo.Client)
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	mongo, err := pm.ConnectToMongoDB(a.ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
	}
	mongoClient = mongo
}

// Expose the CheckKey function
func (a *App) CheckKey(key string) bool {
	return pm.CheckKey(key, mongoClient)
}

func (a *App) GetByPlat(plat string) []string {
	return pm.RetrieveByPlatform(plat, mongoClient)
}

func (a *App) PrintAllRecords(plat string) [][]string {
	return pm.PrintAllRecords(mongoClient)
}

func (a *App) InsertOwn(plat, us, pass string) error {
	user := pm.User{Platform: plat, Username: us, Password: pm.Encrypt(pass)}
	return pm.InsertDocument(mongoClient, user)
}

func (a *App) InsertDocument(plat, us string) error {
	pass, err := pm.GenPassword(20)
	if err != nil {
		return err
	}
	user := pm.User{Platform: plat, Username: us, Password: pm.Encrypt(pass)}
	return pm.InsertDocument(mongoClient, user)
}

func (a *App) DeleteByPlatform(plat string) {
	pm.DeleteByPlatform(plat, mongoClient)
}

// Expose a struct containing functions to be called from the frontend
func (a *App) GetFrontendFunctions() *FrontendFunctions {
	return &FrontendFunctions{App: a}
}

// FrontendFunctions struct
type FrontendFunctions struct {
	App *App
}
