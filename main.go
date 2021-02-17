package main

import (
	"fmt"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(NewApp)

	err := container.Invoke(func(app *App) {
		app.Start()
	})

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
	}
}
