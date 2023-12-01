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

// Expose a struct containing functions to be called from the frontend
func (a *App) GetFrontendFunctions() *FrontendFunctions {
	return &FrontendFunctions{App: a}
}

// FrontendFunctions struct
type FrontendFunctions struct {
	App *App
}
