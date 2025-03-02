package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goBlog/pkg/di"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print(err.Error())
	}
}

func main() {
	di.ConfigureApp().Run()
}
