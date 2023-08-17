package transport

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/dkhrunov/hexagonal-architecture/internal/common/server"
	"github.com/dkhrunov/hexagonal-architecture/internal/domain"
)

// PortService is a port service
type PortService interface {
	GetPort(ctx context.Context, id string) (*domain.Port, error)
	CountPorts(ctx context.Context) (int, error)
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
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
		if errors.Is(err, domain.ErrorNotFound) {
			server.NotFound("port not found", err, w, r)
			return
		}
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

// CountPorts returns total ports stored in DB
func (h HttpServer) CountPorts(w http.ResponseWriter, r *http.Request) {
	total, err := h.service.CountPorts(r.Context())
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOk(map[string]int{"total": total}, w, r)
}

func (h HttpServer) UploadPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("uploading ports...")

	portChan := make(chan PortDto)
	errChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		err := readPorts(r.Context(), r.Body, portChan)
		if err != nil {
			errChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()
	portCounter := 0
	for {
		select {
		case <-r.Context().Done():
			log.Printf("request cancelled")
			return
		case <-doneChan:
			log.Printf("finished reading ports")
			server.RespondOk(map[string]int{"count": portCounter}, w, r)
			return
		case err := <-errChan:
			log.Printf("error while parsing port json: %+v", err)
			server.BadRequest("invalid json", err, w, r)
			return
		case port := <-portChan:
			portCounter++
			log.Printf("[%d] received port: %+v", portCounter, port)
			p, err := portDtoToDomain(&port)
			if err != nil {
				server.BadRequest("can not convert dto port to domain", err, w, r)
				return
			}
			if err := h.service.CreateOrUpdatePort(r.Context(), p); err != nil {
				server.RespondWithError(err, w, r)
				return
			}
		}
	}
}
