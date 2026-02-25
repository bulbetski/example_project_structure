package deps

import (
	"github.com/samber/do"

	rpchealth "github.com/builbetski/example_project_structure/internal/rpctransport/health"
)

func (c *Container) GetHealthGRPCServer() *rpchealth.Server {
	return do.MustInvoke[*rpchealth.Server](c.i)
}

func (c *Container) provideServers() {
	do.Provide(c.i, func(*do.Injector) (*rpchealth.Server, error) {
		svc := c.GetHealthService()
		return rpchealth.NewServer(svc), nil
	})
}
