package inmem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
	"github.com/dkhrunov/hexagonal-architecture/internal/repository"
)

type PortInmemStore struct {
	data map[string]*repository.PortEntity
	mu   sync.RWMutex
}

func NewPortInmemStore() *PortInmemStore {
	return &PortInmemStore{
		data: make(map[string]*repository.PortEntity),
	}
}

func (s *PortInmemStore) GetPort(_ context.Context, id string) (*domain.Port, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	storePort, exists := s.data[id]
	if !exists {
		return nil, domain.ErrorNotFound
	}

	domainPort, err := repository.PortEntityToDomain(storePort)
	if err != nil {
		return nil, fmt.Errorf("PortEntityToDomain failed: %w", err)
	}

	return domainPort, nil
}

func (s *PortInmemStore) CountPorts(_ context.Context) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data), nil
}

func (s *PortInmemStore) CreateOrUpdatePort(ctx context.Context, p *domain.Port) error {
	if p == nil {
		return domain.ErrorNil
	}

	storePort, err := repository.PortDomainToEntity(p)
	if err != nil {
		return fmt.Errorf("PortDomainToEntity failed: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[storePort.ID]
	if exists {
		return s.updatePort(ctx, storePort)
	}

	return s.createPort(ctx, storePort)
}

func (s *PortInmemStore) createPort(_ context.Context, p *repository.PortEntity) error {
	if p == nil {
		return domain.ErrorNil
	}

	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt

	s.data[p.ID] = p

	return nil
}

func (s *PortInmemStore) updatePort(_ context.Context, p *repository.PortEntity) error {
	if p == nil {
		return domain.ErrorNil
	}

	storePort, exists := s.data[p.ID]
	if !exists {
		return domain.ErrorNotFound
	}

	storePortCopy := storePort.Copy()

	storePortCopy.Name = p.Name
	storePortCopy.Code = p.Code
	storePortCopy.City = p.City
	storePortCopy.Country = p.Country
	storePortCopy.Alias = append([]string(nil), p.Alias...)
	storePortCopy.Regions = append([]string(nil), p.Regions...)
	storePortCopy.Coordinates = [2]float64{p.Coordinates[0], p.Coordinates[1]}
	storePortCopy.Province = p.Province
	storePortCopy.Timezone = p.Timezone
	storePortCopy.Unlocs = append([]string(nil), p.Unlocs...)

	storePortCopy.UpdatedAt = time.Now()

	s.data[p.ID] = storePortCopy

	return nil
}
