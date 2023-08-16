package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

func portDtoToDomain(p *PortDto) (*domain.Port, error) {
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

// readPorts reads ports from provided reader and sends them to portChan
func readPorts(ctx context.Context, r io.Reader, portChan chan<- PortDto) error {
	decoder := json.NewDecoder(r)

	// Read opening delimiter
	openingToken, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening delimited: %w", err)
	}

	// Make sure opening delimiter is `{`
	if openingToken != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", openingToken)
	}

	for decoder.More() {
		// Check if context is cancelled
		if ctx.Err() != nil {
			return ctx.Err()
		}
		// Read the port ID
		t, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read port ID: %w", err)
		}
		// Make sure port ID is a string
		portID, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", t)
		}

		// Read the port and send it to the channel
		var portDto PortDto
		if err := decoder.Decode(&portDto); err != nil {
			return fmt.Errorf("failed to decode port: %w", err)
		}

		portDto.ID = portID
		portChan <- portDto
	}

	// Read closing delimiter
	closingToken, err := decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read closing delimited: %w", err)
	}

	// Make sure closing delimiter is `}`
	if closingToken != json.Delim('}') {
		return fmt.Errorf("expected }, got %v", closingToken)
	}

	return nil
}
