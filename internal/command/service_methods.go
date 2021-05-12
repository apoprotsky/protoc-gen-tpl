package command

import (
	"os"
)

// Execute command
func (svc *Service) Execute() {
	if err := svc.rootCommand.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// NeedHelp flag
func (svc *Service) NeedHelp() bool {
	return svc.help
}
