package memory

import (
	"context"
	"errors"
	"sync"

	"{{cookiecutter.gomodule_uri}}/pkg/domain"

	"github.com/rs/zerolog"
)

type DemoRepository struct {
	m sync.Map
}

// Store mock for in-memory storage
func (r *DemoRepository) Store(d domain.Demo) error {
	r.m.Store(d.ID, d)
	return nil
}

// GetByID mock for in-memory storage.
func (r *DemoRepository) LoadByID(ID string) (*domain.Demo, error) {
	v, ok := r.m.Load(ID)
	if !ok {
		// Match upper/db4 db.ErrNoMoreRows
		return nil, errors.New(`upper: no more rows in this result set`)
	}

	d := v.(domain.Demo)
	return &d, nil
}

// NewDemoRepository instantiates an example memory repository
func NewDemoRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*DemoRepository, error) {
	return &DemoRepository{}, nil
}
