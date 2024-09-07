package app

import (
	"errors"
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
	"github.com/dittonetwork/executor-avs/pkg/service"
)

const (
	appName                  = "AVS-operator"
	version                  = "1.0.0"
	defaultShutdownTimeout   = 30 * time.Second
	defaultSimulationTimeout = 5 * time.Second
	devaultIgnoreEndBlocks   = 3
)

var (
	env, addr, diagnosticsAddr string
	shutdownTimeout            time.Duration
	autoDeactivate             bool
	ignoreEndBlocks            uint
	simulationTimeout          time.Duration
)

func initRunFlags(cmd *cobra.Command) {
	// TODO: handle these flags with cobra as well
	cmd.Flags().StringVar(&env, "env", "dev", "Operator environment")
	cmd.Flags().StringVar(&addr, "addr", ":8080", "Operator addr")
	cmd.Flags().StringVar(&diagnosticsAddr, "diagnostics-addr", ":7070", "Operator diagnostics addr")
	cmd.Flags().DurationVar(&shutdownTimeout, "shutdown-timeout", defaultShutdownTimeout, "Graceful shutdown timeout")
	cmd.Flags().DurationVar(
		&simulationTimeout,
		"simulation-timeout",
		defaultSimulationTimeout,
		"RPC simulation execution timeout",
	)
	cmd.Flags().BoolVar(&autoDeactivate, "auto-deactivate", true, "Deactivate the operator on shutdown")
	cmd.Flags().UintVar(
		&ignoreEndBlocks,
		"ignore-end-block",
		devaultIgnoreEndBlocks,
		"Block amount to ignore execution at the end of epoch",
	)
}

func Run(cfg *Config) (*sync.WaitGroup, error) {
	if len(cfg.ExecutorPrivateKey) == 0 {
		return nil, errors.New("EXECUTOR_PRIVATE_KEY env in not set")
	}

	service.Init(appName, env, service.WithDiagnosticsServer(diagnosticsAddr))

	conn, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		return nil, err
	}

	// adapters
	ethClient, err := ethereum.NewClient(conn, cfg.ContractAddress, cfg.ExecutorPrivateKey)
	// we keep this client opened until the end, it closes implicitly
	if err != nil {
		return nil, err
	}

	entryPoint, err := dittoentrypoint.New(conn, cfg.ContractAddress, cfg.ExecutorPrivateKey)
	if err != nil {
		return nil, err
	}

	// services
	executorService := executor.NewService(
		executor.NewExecutor(
			ethClient,
			entryPoint,
			big.NewInt(int64(ignoreEndBlocks)),
			simulationTimeout,
			executor.WithMetrics(),
			executor.WithCustomLiveCycle(autoDeactivate),
		),
	)

	wg := service.RunWait(
		executorService,
		rest.NewServer(addr, shutdownTimeout,
			endpoint.NewHealthCheckEndpoint(),
			endpoint.NewInformationEndpoint(appName, version),
			endpoint.NewServicesEndpoint(),
			endpoint.NewServiceHealthEndpoint(),
		),
	)

	return wg, nil
}
