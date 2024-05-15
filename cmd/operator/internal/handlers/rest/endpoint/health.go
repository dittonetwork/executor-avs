package endpoint

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthCheckEndpoint struct {
}

func NewHealthCheckEndpoint() *HealthCheckEndpoint {
	return &HealthCheckEndpoint{}
}

func (h *HealthCheckEndpoint) Setup(router *httprouter.Router) {
	router.GET("/eigen/node/health", h.handle)
}

func (h *HealthCheckEndpoint) handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
