package demo

import (
	"net/http"

	"{{cookiecutter.gomodule_uri}}/pkg/domain"
	"{{cookiecutter.gomodule_uri}}/pkg/service"

	"github.com/go-chi/render"
)

// FetchResponse is the shape of data for a loaded demo record
type FetchResponse struct {
	domain.Demo
}

// Render satisfies the chi interface
func (fr *FetchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// CreateResponse contains the ID post demo creation
type CreateResponse struct {
	ID string
}

// Render setups up the correct http status code.
func (cr *CreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d domain.Demo, _ *service.DemoService) *FetchResponse {
	return &FetchResponse{d}
}

// NewCreateResponse instantiates a new response when demo is created
func NewCreateResponse(d domain.Demo, _ *service.DemoService) *CreateResponse {
	return &CreateResponse{ID: "demo"}
}
