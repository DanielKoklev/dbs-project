package nbdtarget

import (
	"context"
	"errors"
	"net"
	"sync/atomic"
)

type Backend interface {
	ReadAt(ctx context.Context, p []byte, off int64) (int, error)
	WriteAt(ctx context.Context, p []byte, off int64) (int, error)
	Flush(ctx context.Context) error
	Size(ctx context.Context) (int64, error)
}

type Server struct {
	BE Backend
	Conns atomic.Int64
}

func (s *Server) Serve(l net.Listener) error {
	// Implement nbd protocol handshake
	return errors.New("nbd target not implemented")
}

// Mock backend for dev
func NewMockBackend(size int64) Backend { return &mock{size: size} }

type mock struct { size int64 }

func (m *mock) ReadAt(ctx context.Context, p []byte, off int64) (int, error) { return len(p), nil }
func (m *mock) WriteAt(ctx context.Context, p []byte, off int64) (int, error) { return len(p), nil }
func (m *mock) Flush(ctx context.Context) error { return nil }
func (m *mock) Size(ctx context.Context) (int64, error) { return m.size, nil }