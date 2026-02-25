package deps

import (
	"github.com/samber/do"

	httphealth "github.com/builbetski/example_project_structure/internal/httptransport/health"
)

func (c *Container) GetHealthHTTPHandler() *httphealth.Handler {
	return do.MustInvoke[*httphealth.Handler](c.i)
}

func (c *Container) provideHandlers() {
	do.Provide(c.i, func(*do.Injector) (*httphealth.Handler, error) {
		svc := c.GetHealthService()
		return httphealth.NewHandler(svc), nil
	})
}
