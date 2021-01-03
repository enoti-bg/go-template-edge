package service

import (
	"{{cookiecutter.gomodule_uri}}/pkg/domain"

	"github.com/rs/zerolog"
)

// DemoService contains the business logic for demos
type DemoService struct {
	Repository domain.DemoRepository
	Logger     *zerolog.Logger
}

// GetByID - Proxy to repository
func (ds DemoService) GetByID(ID string) (*domain.Demo, error) {
	// Stuff happens
	return ds.Repository.LoadByID(ID)
}

// Store converts a raw label to a demo record and saves to the respective repository.
func (ds DemoService) Store(label string) error {
	// Stuff happens
	return ds.Repository.Store(domain.Demo{ID: "demo", Label: label})
}
