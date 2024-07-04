package main

import (
	"awesomeProject/pkg/di"
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err.Error())
	}
}

func main() {
	di.ConfigureApp().Run()
}
