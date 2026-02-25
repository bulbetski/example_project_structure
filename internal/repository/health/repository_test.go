package health

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakePinger struct {
	err    error
	called bool
}

func (f *fakePinger) Ping(_ context.Context) error {
	f.called = true
	return f.err
}

func TestRepository_Ping_OK(t *testing.T) {
	pinger := &fakePinger{}
	repo := NewRepository(pinger)

	err := repo.Ping(t.Context())

	require.NoError(t, err)
	require.True(t, pinger.called)
}

func TestRepository_Ping_Error(t *testing.T) {
	pinger := &fakePinger{err: errors.New("boom")}
	repo := NewRepository(pinger)

	err := repo.Ping(t.Context())

	require.Error(t, err)
	require.True(t, pinger.called)
}
