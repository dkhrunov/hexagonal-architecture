package services

import (
	"context"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

// PortRepository is a port repository for the port service
type PortRepository interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	CountPorts(ctx context.Context) (int, error)
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
}

// PortService is a port service
type PortService struct {
	repository PortRepository
}

// NewPortService creates a new port service
func NewPortService(repository PortRepository) PortService {
	return PortService{
		repository: repository,
	}
}

// GetPort retuns a port by id
func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	return s.repository.GetPort(ctx, id)
}

// CountPorts returns the number of ports
func (s PortService) CountPorts(ctx context.Context) (int, error) {
	return s.repository.CountPorts(ctx)
}

// CreateOrUpdatePort creates or updates a port
func (s PortService) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	return s.repository.CreateOrUpdatePort(ctx, port)
}
