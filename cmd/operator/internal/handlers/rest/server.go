package rest

import (
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/middleware"
	"github.com/dittonetwork/executor-avs/pkg/service"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"time"
)

type Endpoint interface {
	Setup(router *httprouter.Router)
}

func NewServer(addr string, shutdownTimeout time.Duration,
	endpoints ...Endpoint,
) *service.HTTPServer {
	router := httprouter.New()

	for _, api := range endpoints {
		api.Setup(router)
	}

	return service.NewHTTPServer(
		addr,
		shutdownTimeout,
		alice.New(
			middleware.Recover(log.Default(), middleware.LogPanicRequest),
			middleware.Tracing(),
		).Then(router),
	)
}
