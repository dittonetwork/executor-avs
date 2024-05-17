package app

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/node/ethereum"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/handlers/rest"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/handlers/rest/endpoint"
	"github.com/dittonetwork/executor-avs/cmd/operator/internal/services/executor"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/dittonetwork/executor-avs/pkg/service"
)

const (
	appName                = "AVS-operator"
	version                = "1.0.0"
	defaultShutdownTimeout = 30 * time.Second
)

var (
	env, addr, diagnosticsAddr string
	shutdownTimeout            time.Duration
)

func initRunFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&env, "env", "dev", "Operator environment")
	cmd.Flags().StringVar(&addr, "addr", ":8080", "Operator addr")
	cmd.Flags().StringVar(&diagnosticsAddr, "diagnostics-addr", ":7070", "Operator diagnostics addr")
	cmd.Flags().DurationVar(&shutdownTimeout, "shutdown-timeout", defaultShutdownTimeout, "Graceful shutdown timeout")
}

func Run() *sync.WaitGroup {
	service.Init(appName, env, service.WithDiagnosticsServer(diagnosticsAddr))

	// adapters
	conn, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.With(log.Err(err)).Fatal("ethereum client init error")
	}

	ethClient := ethereum.NewClient(conn)

	entryPoint, err := dittoentrypoint.New(conn, contractAddress, privateKey)
	if err != nil {
		log.With(log.Err(err)).Fatal("dittoentrypoint init error")
	}

	// services
	executorService := executor.NewService(ethClient, entryPoint, contractAddress)

	return service.RunWait(
		executorService,
		rest.NewServer(addr, shutdownTimeout,
			endpoint.NewHealthCheckEndpoint(),
			endpoint.NewInformationEndpoint(appName, version),
			endpoint.NewServicesEndpoint(),
			endpoint.NewServiceHealthEndpoint(),
		),
	)
}
