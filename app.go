package main

import "fmt"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	fmt.Println("Start application...")
	fmt.Println("Exit...")
}
