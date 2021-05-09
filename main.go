package main

import (
	"github.com/apoprotsky/proto-gen-tpl/internal/command"
)

func main() {
	commandService := command.GetService()
	commandService.Execute()
}
