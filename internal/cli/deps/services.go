package deps

import (
	"github.com/samber/do"

	svchealth "github.com/builbetski/example_project_structure/internal/service/health"
)

func (c *Container) GetHealthService() *svchealth.Service {
	return do.MustInvoke[*svchealth.Service](c.i)
}

func (c *Container) provideServices() {
	do.Provide(c.i, func(*do.Injector) (*svchealth.Service, error) {
		repo := c.GetHealthRepository()
		return svchealth.NewService(repo), nil
	})
}
