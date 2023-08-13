package transport

import (
	"context"
	"net/http"

	"github.com/dkhrunov/hexagonal-architecture/internal/common/server"
	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

// PortService is a port service
type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}

// HttpServer is a HTTP server for ports
type HttpServer struct {
	service PortService
}

// NewHttpServer creates a new HTTP server for ports
func NewHttpServer(service PortService) HttpServer {
	return HttpServer{
		service: service,
	}
}

func (h HttpServer) GetPort(w http.ResponseWriter, r *http.Request) {
	port, err := h.service.GetPort(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := PortDto{
		ID:          port.ID(),
		Name:        port.Name(),
		City:        port.City(),
		Country:     port.Country(),
		Alias:       port.Alias(),
		Regions:     port.Regions(),
		Coordinates: port.Coordinates(),
		Province:    port.Province(),
		Timezone:    port.Timezone(),
		Unlocs:      port.Unlocs(),
	}

	server.RespondOk(response, w, r)
}
