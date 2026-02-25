package health

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

type fakeService struct {
	err error
}

func (f *fakeService) Ping(_ context.Context) error {
	return f.err
}

func TestHandler_Health_OK(t *testing.T) {
	e := echo.New()
	service := &fakeService{}
	h := NewHandler(service)
	h.RegisterRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	require.Contains(t, rec.Body.String(), "\"status\":\"ok\"")
}

func TestHandler_Health_Error(t *testing.T) {
	e := echo.New()
	service := &fakeService{err: errors.New("boom")}
	h := NewHandler(service)
	h.RegisterRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusServiceUnavailable, rec.Code)
	require.Contains(t, rec.Body.String(), "\"error\":{\"message\":\"boom\"}")
}
