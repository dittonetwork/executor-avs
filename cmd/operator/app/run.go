package app

import (
	"flag"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/handlers/rest"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/handlers/rest/endpoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/service"
	"os"
	"sync"
	"time"
)

const (
	appName                = "AVS-operator"
	version                = "1.0.0"
	defaultShutdownTimeout = 30 * time.Second
)

var (
	fs = flag.NewFlagSet(appName, flag.ExitOnError)

	env             = fs.String("env", "dev", "Operator environment")
	addr            = fs.String("addr", ":8080", "Operator addr")
	diagnosticsAddr = fs.String("diagnostics-addr", ":7070", "Operator diagnostics addr")
	shutdownTimeout = fs.Duration("shutdown-timeout", defaultShutdownTimeout, "Graceful shutdown timeout")
)

func Run() *sync.WaitGroup {
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.With(log.Err(err)).Fatal("parse flags error")
	}

	service.Init(appName, *env, service.WithDiagnosticsServer(*diagnosticsAddr))

	executorService := executor.NewService()

	return service.RunWait(
		executorService,
		rest.NewServer(*addr, *shutdownTimeout,
			endpoint.NewHealthCheckEndpoint(),
			endpoint.NewInformationEndpoint(appName, version),
			endpoint.NewServicesEndpoint(),
			endpoint.NewServiceHealthEndpoint(),
		),
	)
}
