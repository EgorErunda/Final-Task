package main

import (
	"github.com/EgorErunda/FinalTaskSprint_1/internal/application"
)

func main() {
	application := application.New()
	application.RunServer()
}
