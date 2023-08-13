package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dkhrunov/hexagonal-architecture/internal/common/errors"
	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
	"github.com/google/uuid"
)

// PortService is a port service
type PortService struct{}

// NewPortService creates a new port service
func NewPortService() PortService {
	return PortService{}
}

// GetPort retuns a port by id
func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	if id == "" {
		return nil, errors.NewIncorrectInputError(
			"Query param 'id' not provided",
			"Request URL not contains 'id' query param. This parameter is required for this request",
		)
	}
	if idInt, _ := strconv.Atoi(id); idInt > 5 {
		return nil, errors.NewNotFoundError(
			"Port not found",
			fmt.Sprintf("Port with id:%v not found", id),
		)
	}
	randomID := uuid.New().String()
	return domain.NewPort(randomID, randomID, randomID, randomID, randomID, []string{randomID}, []string{randomID}, [2]float64{1.0, 2.0}, randomID, randomID, nil)
}
