package deps

import (
	"github.com/samber/do"

	repohealth "github.com/builbetski/example_project_structure/internal/repository/health"
)

func (c *Container) GetHealthRepository() *repohealth.Repository {
	return do.MustInvoke[*repohealth.Repository](c.i)
}

func (c *Container) provideRepositories() {
	do.Provide(c.i, func(*do.Injector) (*repohealth.Repository, error) {
		pool := c.GetDatabase()
		return repohealth.NewRepository(pool), nil
	})
}
