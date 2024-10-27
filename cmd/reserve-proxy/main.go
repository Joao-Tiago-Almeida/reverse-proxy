package main

import (
	"fmt"
)

// Define the ports for the application
type App struct {
	db database
}

func (App) Run() {
	// Run the application
	fmt.Println("Running the application")
}

func main() {
	// Create a new instance of the application
	app := App{
		db: initDatabase(),
	}

	app.Run()
}
