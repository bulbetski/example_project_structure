package health

import "context"

type Pinger interface {
	Ping(ctx context.Context) error
}

type Repository struct {
	pinger Pinger
}

func NewRepository(pinger Pinger) *Repository {
	return &Repository{pinger: pinger}
}

func (r *Repository) Ping(ctx context.Context) error {
	return r.pinger.Ping(ctx)
}
