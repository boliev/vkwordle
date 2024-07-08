package vkwordle

import "fmt"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	fmt.Println("vkwordle is running")
}
