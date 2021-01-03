package demo

import (
	"errors"
	"net/http"

	"{{cookiecutter.gomodule_uri}}/pkg/domain"
	"{{cookiecutter.gomodule_uri}}/pkg/service"
	weberror "{{cookiecutter.gomodule_uri}}/pkg/web/error"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

const (
	// You decide if you want to wrap errors or
	// will use values.
	ErrGetDemoParam    = "get_demo_param"
	ErrGetDemoLoad     = "get_demo_load"
	ErrCreateDemoParam = "create_demo_param"
	ErrCreateDemoStore = "create_demo_store"
)

// Handler is just a route collection
type Handler struct{}

// GetDemo Load a specific demo by ID - only "demo" will be found
func (h Handler) GetDemo(logger *zerolog.Logger, ds *service.DemoService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		if ID == "" {
			render.Render(w, r, weberror.NewErrorResponse(ErrGetDemoParam, http.StatusBadRequest, errors.New("passed an empty ID"), logger))
			return
		}

		demo, err := ds.GetByID(ID)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrGetDemoLoad, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewFetchResponse(*demo, ds))
	}
}

// CreateDemo allows HTTP creation.
func (h Handler) CreateDemo(logger *zerolog.Logger, ds *service.DemoService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrCreateDemoParam, http.StatusBadRequest, err, logger))
			return
		}

		err := ds.Store(request.Label)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrCreateDemoStore, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewCreateResponse(domain.Demo{ID: "demo", Label: request.Label}, ds))
	}
}

// Routes for demo create/read
func (h Handler) Routes(logger *zerolog.Logger, ds *service.DemoService) chi.Router {
	r := chi.NewRouter()

	r.Post("/demo", h.CreateDemo(logger, ds))
	r.Get("/demo/{ID}", h.GetDemo(logger, ds))

	return r
}
