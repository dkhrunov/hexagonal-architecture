package domain

import "fmt"

type Port struct {
	id          string
	name        string
	code        string
	city        string
	country     string
	alias       []string
	regions     []string
	coordinates [2]float64
	province    string
	timezone    string
	unlocs      []string
}

// NewPort creates a new port
func NewPort(id, name, code, city, country string, alias, regions []string, coordinates [2]float64, province, timezone string, unlocs []string) (*Port, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: port id is required", ErrorRequired)
	}
	if name == "" {
		return nil, fmt.Errorf("%w: port name is required", ErrorRequired)
	}
	if city == "" {
		return nil, fmt.Errorf("%w: port city is required", ErrorRequired)
	}
	if country == "" {
		return nil, fmt.Errorf("%w: port country is required", ErrorRequired)
	}

	return &Port{
		id:          id,
		name:        name,
		code:        code,
		city:        city,
		country:     country,
		alias:       alias,
		regions:     regions,
		coordinates: coordinates,
		province:    province,
		timezone:    timezone,
		unlocs:      unlocs,
	}, nil
}

// ID returns the port id
func (p *Port) ID() string {
	return p.id
}

// Name returns the port name
func (p *Port) Name() string {
	return p.name
}

// SetName sets the port name
func (p *Port) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("%w: port name is required", ErrorRequired)
	}
	p.name = name
	return nil
}

// Code returns the port code.
func (p *Port) Code() string {
	return p.code
}

// City returns the port city.
func (p *Port) City() string {
	return p.city
}

// Country returns the port country.
func (p *Port) Country() string {
	return p.country
}

// Alias returns the port alias.
func (p *Port) Alias() []string {
	return p.alias
}

// Regions returns the port regions.
func (p *Port) Regions() []string {
	return p.regions
}

// Coordinates returns the port coordinates.
func (p *Port) Coordinates() [2]float64 {
	return p.coordinates
}

// Province returns the port province.
func (p *Port) Province() string {
	return p.province
}

// Timezone returns the port timezone.
func (p *Port) Timezone() string {
	return p.timezone
}

// Unlocs returns the port unlocs.
func (p *Port) Unlocs() []string {
	return p.unlocs
}
