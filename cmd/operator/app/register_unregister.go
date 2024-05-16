package app

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/pkg/log"
)

func InitDittoEntryPoint(nodeURL, dittoContractAddress string) (*dittoentrypoint.DittoEntryPoint, error) {
	conn, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, fmt.Errorf("ethereum client dial: %w", err)
	}

	entryPoint, err := dittoentrypoint.New(conn, dittoContractAddress)
	if err != nil {
		return nil, fmt.Errorf("dittoentrypoint client init: %w", err)
	}

	return entryPoint, nil
}

func RegisterExecutor(nodeURL, dittoContractAddress string) {
	dep, err := InitDittoEntryPoint(nodeURL, dittoContractAddress)
	if err != nil {
		log.With(log.Err(err)).Fatal("ditto entrypoint init error")
	}

	if err = dep.RegisterExecutor(context.Background()); err != nil {
		log.With(log.Err(err)).Fatal("register executor error")
	}
}

func UnregisterExecutor(nodeURL, dittoContractAddress string) {
	dep, err := InitDittoEntryPoint(nodeURL, dittoContractAddress)
	if err != nil {
		log.With(log.Err(err)).Fatal("ditto entrypoint init error")
	}

	if err = dep.UnregisterExecutor(context.Background()); err != nil {
		log.With(log.Err(err)).Fatal("unregister executor error")
	}
}
