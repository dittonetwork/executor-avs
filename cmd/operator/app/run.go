package app

import (
	"math/big"
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
	devaultIgnoreEndBlocks = 3
)

var (
	env, addr, diagnosticsAddr string
	shutdownTimeout            time.Duration
	autoDeactivate             bool
	ignoreEndBlocks            uint
)

func initRunFlags(cmd *cobra.Command) {
	// TODO: handle these flags with cobra as well
	cmd.Flags().StringVar(&env, "env", "dev", "Operator environment")
	cmd.Flags().StringVar(&addr, "addr", ":8080", "Operator addr")
	cmd.Flags().StringVar(&diagnosticsAddr, "diagnostics-addr", ":7070", "Operator diagnostics addr")
	cmd.Flags().DurationVar(&shutdownTimeout, "shutdown-timeout", defaultShutdownTimeout, "Graceful shutdown timeout")
	cmd.Flags().BoolVar(&autoDeactivate, "auto-deactivate", true, "Do not deactivate the operator on shutdown")
	cmd.Flags().UintVar(&ignoreEndBlocks, "ignore-end-block", devaultIgnoreEndBlocks, "Graceful shutdown timeout")
}

func Run(cfg *CommonFlags) *sync.WaitGroup {
	service.Init(appName, env, service.WithDiagnosticsServer(diagnosticsAddr))

	conn, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		log.With(log.Err(err)).Fatal("ether client dial error")
	}

	// adapters
	ethClient, err := ethereum.NewClient(conn, cfg.ContractAddress, cfg.PrivateKey)
	if err != nil {
		log.With(log.Err(err)).Fatal("ethereum client init error")
	}

	entryPoint, err := dittoentrypoint.New(conn, cfg.ContractAddress, cfg.PrivateKey)
	if err != nil {
		log.With(log.Err(err)).Fatal("dittoentrypoint init error")
	}
	// services
	// TODO: refactor WithMetrics passing
	executorService := executor.NewService(
		executor.NewExecutor(
			ethClient,
			entryPoint,
			big.NewInt(int64(ignoreEndBlocks)),
			executor.WithMetrics(),
			executor.WithCustomLiveCycle(autoDeactivate),
		),
	)

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
