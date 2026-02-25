package health

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeRepo struct {
	err    error
	called bool
}

func (f *fakeRepo) Ping(_ context.Context) error {
	f.called = true
	return f.err
}

func TestService_Ping_OK(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewService(repo)

	err := svc.Ping(t.Context())

	require.NoError(t, err)
	require.True(t, repo.called)
}

func TestService_Ping_Error(t *testing.T) {
	repo := &fakeRepo{err: errors.New("boom")}
	svc := NewService(repo)

	err := svc.Ping(t.Context())

	require.Error(t, err)
	require.True(t, repo.called)
}
