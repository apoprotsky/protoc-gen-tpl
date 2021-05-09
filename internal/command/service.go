package command

import (
	"os"
	"path/filepath"

	"github.com/apoprotsky/protoc-gen-tpl/internal/app"

	gs "github.com/apoprotsky/go-services"
	"github.com/spf13/cobra"
)

// Service struct
type Service struct {
	rootCommand *cobra.Command
	help        bool
}

// GoService initializes service
func (svc *Service) GoService() {
	svc.rootCommand = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: app.Name,
		Run:   rootCommand,
	}
	svc.rootCommand.PersistentFlags().BoolVarP(&svc.help, "help", "h", false, "help message")
}

// GetService returns instance of service
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
