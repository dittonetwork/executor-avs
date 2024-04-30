package endpoint

import (
	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Service interface {
	GetName() string
	GetID() string
	GetDescription() string
	GetStatus() api.ServiceStatusType
}

type ServicesEndpoint struct {
	services []Service

	responder
}

func NewServicesEndpoint(services ...Service) *ServicesEndpoint {
	return &ServicesEndpoint{
		services: services,
	}
}

func (s *ServicesEndpoint) Setup(router *httprouter.Router) {
	router.GET("/eigen/node/services", s.handle)
}

func (s *ServicesEndpoint) handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := api.EigenNodeServicesResponse{
		Services: make([]api.EigenNodeService, 0, len(s.services)),
	}

	for _, service := range s.services {
		response.Services = append(response.Services, api.EigenNodeService{
			ID:          service.GetID(),
			Name:        service.GetName(),
			Description: service.GetDescription(),
			Status:      service.GetStatus(),
		})
	}

	s.RespondEasyJson(r.Context(), w, response)
}
