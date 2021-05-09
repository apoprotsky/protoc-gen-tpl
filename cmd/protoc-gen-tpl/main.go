package main

import (
	"github.com/apoprotsky/prototpl/internal/command"
)

func main() {
	commandService := command.GetService()
	commandService.Execute()
}
