package deps

import (
	"github.com/samber/do"

	"github.com/builbetski/example_project_structure/internal/config"
)

func (c *Container) GetConfig() *config.Config {
	return do.MustInvoke[*config.Config](c.i)
}

func (c *Container) provideConfig() {
	do.Provide(c.i, func(*do.Injector) (*config.Config, error) {
		return config.Load()
	})
}
