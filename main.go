package main

import (
	"github.com/apoprotsky/protoc-gen-tpl/internal/command"
)

func main() {
	commandService := command.GetService()
	commandService.Execute()
}
