package endpoint

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	api "github.com/dittonetwork/executor-avs/api/operator"
)

type InformationEndpoint struct {
	NodeName string
	Version  string

	responder
}

func NewInformationEndpoint(name string, version string) *InformationEndpoint {
	return &InformationEndpoint{
		NodeName: name,
		Version:  version,
	}
}

func (i *InformationEndpoint) Setup(router *httprouter.Router) {
	router.GET("/eigen/node", i.handle)
}

func (i *InformationEndpoint) handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	i.RespondEasyJSON(r.Context(), w, api.EigenNodeResponse{
		NodeName:    i.NodeName,
		NodeVersion: i.Version,
		SpecVersion: "0.0.1",
	})
}
