package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Define the ports for the application
type App struct {
	db database
}

func (app App) Run() {
	// Run the application
	fmt.Println("Running the application")
	values, _ := app.db.Find(map[string]string{"host": "iris"})
	fmt.Println("Data", values)
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Create a new instance of the application
	app := App{
		db: initDatabase(),
	}

	app.Run()
}
