package generator

import (
	gs "github.com/apoprotsky/go-services"
)

// Service struct
type Service struct {
	golang     bool
	typescript bool
	php        bool
}

// GoService initializes service
func (svc *Service) GoService() {
}

// GetService returns instance of service
func GetService() *Service {
	return gs.Get((*Service)(nil)).(*Service)
}
