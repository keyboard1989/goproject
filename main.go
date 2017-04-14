package main

import "github.com/keyboard1989/goproject/application"

func main() {
	app := application.NewApp()
	service := application.NewService(app)

	service.Run()
}
