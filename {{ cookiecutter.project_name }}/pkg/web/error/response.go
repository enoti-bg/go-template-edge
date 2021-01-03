package response

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

// ErrorResponse structure shared by all handlers
type ErrorResponse struct {
	HTTPStatusCode  int    `json:"-"`
	Status          int    `json:"status"`
	Message         string `json:"message"`
	InternalMessage string `json:"-"`
}

// Render implementation satisfying render.Renderer
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// NewErrorResponse logs the error and returns renderable response.
func NewErrorResponse(msg string, status int, err error, logger *zerolog.Logger) render.Renderer {
	internalMessage := ""
	if err != nil {
		logger.Error().Err(err).Msg("HTTP error response")
	}
	return &ErrorResponse{
		HTTPStatusCode:  status,
		Status:          status,
		Message:         msg,
		InternalMessage: internalMessage,
	}
}
