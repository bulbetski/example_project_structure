package health

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	apidomain "github.com/builbetski/example_project_structure/internal/domain/api"
)

// HealthService defines the dependency interface for the handler.
type HealthService interface {
	Ping(ctx context.Context) error
}

// Handler provides HTTP endpoints for health checks.
type Handler struct {
	svc HealthService
}

func NewHandler(svc HealthService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/healthz", h.Health)
}

func (h *Handler) Health(c echo.Context) error {
	if err := h.svc.Ping(c.Request().Context()); err != nil {
		return c.JSON(http.StatusServiceUnavailable, apidomain.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
