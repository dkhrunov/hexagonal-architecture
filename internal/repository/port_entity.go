package repository

import (
	"errors"
	"time"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

type PortEntity struct {
	ID          string
	Name        string
	Code        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates [2]float64
	Province    string
	Timezone    string
	Unlocs      []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewPortEntity creates a new port entity
func NewPortEntity(id, name, code, city, country string, alias, regions []string, coordinates [2]float64, province, timezone string, unlocs []string) (*PortEntity, error) {
	return &PortEntity{
		ID:          id,
		Name:        name,
		Code:        code,
		City:        city,
		Country:     country,
		Alias:       alias,
		Regions:     regions,
		Coordinates: coordinates,
		Province:    province,
		Timezone:    timezone,
		Unlocs:      unlocs,
	}, nil
}

func (p *PortEntity) Copy() *PortEntity {
	if p == nil {
		return nil
	}
	return &PortEntity{
		ID:          p.ID,
		Name:        p.Name,
		Code:        p.Code,
		City:        p.City,
		Country:     p.Country,
		Alias:       append([]string(nil), p.Alias...),
		Regions:     append([]string(nil), p.Regions...),
		Coordinates: [2]float64{p.Coordinates[0], p.Coordinates[1]},
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      append([]string(nil), p.Unlocs...),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func PortEntityToDomain(p *PortEntity) (*domain.Port, error) {
	if p == nil {
		return nil, errors.New("entity port is nil")
	}
	return domain.NewPort(
		p.ID,
		p.Name,
		p.Code,
		p.City,
		p.Country,
		append([]string(nil), p.Alias...),
		append([]string(nil), p.Regions...),
		[2]float64{p.Coordinates[0], p.Coordinates[1]},
		p.Province,
		p.Timezone,
		append([]string(nil), p.Unlocs...),
	)
}

func PortDomainToEntity(p *domain.Port) (*PortEntity, error) {
	if p == nil {
		return nil, errors.New("domain port is nil")
	}
	return NewPortEntity(
		p.ID(),
		p.Name(),
		p.Code(),
		p.City(),
		p.Country(),
		append([]string(nil), p.Alias()...),
		append([]string(nil), p.Regions()...),
		[2]float64{p.Coordinates()[0], p.Coordinates()[1]},
		p.Province(),
		p.Timezone(),
		append([]string(nil), p.Unlocs()...),
	)
}
