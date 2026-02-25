package deps

import "github.com/samber/do"

type Container struct {
	i *do.Injector
}

func NewContainer() *Container {
	c := &Container{i: do.New()}

	c.provideConfig()
	c.provideDatabase()
	c.provideRepositories()
	c.provideServices()
	c.provideServers()
	c.provideHandlers()

	return c
}
