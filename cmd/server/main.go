package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/Prateek13727/go-rest/internal/transport/http"
)

//App - the struct which contains things like pointer to our database struct
type App struct{}

//Run sets up our application
func (app *App) Run() error {
	fmt.Println("setting up our App")
	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up the server")
		return err
	}
	return nil
}

// Our main entrypoint for the application
func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
