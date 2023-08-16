package repository

import (
	"context"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

// PortRepository is a port repository for the port service
type IPortRepository interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	CountPorts(ctx context.Context) (int, error)
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
}
