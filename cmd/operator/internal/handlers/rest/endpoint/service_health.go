package endpoint

import (
	api "github.com/dittonetwork/executor-avs/api/operator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ServiceHealthEndpoint struct {
	services map[string]Service
}

func NewServiceHealthEndpoint(services ...Service) *ServiceHealthEndpoint {
	var serviceMap = make(map[string]Service)
	for _, service := range services {
		serviceMap[service.GetID()] = service
	}

	return &ServiceHealthEndpoint{
		services: serviceMap,
	}
}

func (s *ServiceHealthEndpoint) Setup(router *httprouter.Router) {
	router.GET("/eigen/node/services/:service_id/health", s.handle)
}

func (s *ServiceHealthEndpoint) handle(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	if service, ok := s.services[param.ByName("service_id")]; ok {
		if service.GetStatus() == api.ServiceStatusTypeActive {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusPartialContent)
		}
	} else {
		w.WriteHeader(http.StatusVariantAlsoNegotiates)
	}
}
