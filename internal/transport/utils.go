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

type JSONEntry struct {
	Data  PortDto
	Error error
}

type JSONStream struct {
	stream chan JSONEntry
}

func NewJSONStream() JSONStream {
	return JSONStream{
		stream: make(chan JSONEntry),
	}
}

func (s JSONStream) Watch() <-chan JSONEntry {
	return s.stream
}

func (s JSONStream) Start(ctx context.Context, r io.Reader) {
	// defer close(s.stream)

	decoder := json.NewDecoder(r)

	// Read opening delimiter. `[` or `{`
	openingToken, err := decoder.Token()
	if err != nil {
		s.stream <- JSONEntry{Error: fmt.Errorf("failed to read opening delimited: %w", err)}
		return
	}
	// Make sure opening delimiter is `{`
	if openingToken != json.Delim('{') {
		s.stream <- JSONEntry{Error: fmt.Errorf("expected {, got %v", openingToken)}
		return
	}

	// Read file content as long as there is something.
	for decoder.More() {
		// Check if context is cancelled
		if ctx.Err() != nil {
			s.stream <- JSONEntry{Error: ctx.Err()}
			return
		}

		// Read the port ID
		t, err := decoder.Token()
		if err != nil {
			s.stream <- JSONEntry{Error: fmt.Errorf("failed to read port ID: %w", err)}
			return
		}

		// Make sure port ID is a string
		portID, ok := t.(string)
		if !ok {
			s.stream <- JSONEntry{Error: fmt.Errorf("expected string, got %v", t)}
			return
		}

		// Read the port and send it to the channel
		var portDto PortDto
		if err := decoder.Decode(&portDto); err != nil {
			s.stream <- JSONEntry{Error: fmt.Errorf("failed to decode: %w", err)}
			return
		}

		portDto.ID = portID
		s.stream <- JSONEntry{Data: portDto}
	}

	// Read closing delimiter
	closingToken, err := decoder.Token()
	if err != nil {
		s.stream <- JSONEntry{Error: fmt.Errorf("failed to read closing delimited: %w", err)}
		return
	}

	// Make sure closing delimiter is `}`
	if closingToken != json.Delim('}') {
		s.stream <- JSONEntry{Error: fmt.Errorf("expected }, got %v", closingToken)}
		return
	}
}
