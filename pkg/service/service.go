package service

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/dittonetwork/executor-avs/pkg/labels"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

type Service struct {
	appName         string
	diagnosticsAddr string
}

// service singleton.
var svc = Service{}

type Option func(*Service)

func WithDiagnosticsServer(addr string) Option {
	return func(s *Service) {
		s.diagnosticsAddr = addr
	}
}

func Init(appName, env string, opts ...Option) {
	SetAlive(true)

	svc.appName = appName
	for _, opt := range opts {
		opt(&svc)
	}

	labels.Add(map[string]string{"app": appName, "env": env})
	SetInfo(labels.Labels)
	log.SetDefault(log.With(labels.LogFields()...))

	if env == "dev" {
		log.SetDefault(log.NewLogger(log.WithFormatter(log.FormatterText), log.WithLevel(log.DebugLevel)))
	}

	log.Info("initializing app")

	if svc.diagnosticsAddr != "" {
		StartDiagnosticsServer(svc.diagnosticsAddr)
	}
}

type StartStopper interface {
	Start()
	Stop()
}

func RunWait(services ...StartStopper) *sync.WaitGroup {
	log.Info("starting app")
	for _, s := range services {
		s.Start()
	}
	SetReady(true)
	log.Info("app ready")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer wg.Done()
		firstSignal := true
		for sig := range signals {
			log.With(log.String("signal", sig.String())).Info("received signal")
			if firstSignal {
				firstSignal = false
				go func() {
					shutdownServices(services)
					close(signals)
				}()
			} else {
				log.Info("Forcing shutdown due to second signal")
				os.Exit(1) // Force exit if second signal received during shutdown
			}
		}
	}()
	return wg
}

func shutdownServices(services []StartStopper) {
	SetReady(false)
	// stop in reverse order
	for i := len(services) - 1; i >= 0; i-- {
		services[i].Stop()
	}
	log.Info("All services stopped gracefully.")
}
