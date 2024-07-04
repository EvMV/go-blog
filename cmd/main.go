package main

import (
	"awesomeProject/pkg/di"
)

func main() {
	di.BuildContainer().Run()
}
