package app

import (
	"context"
	"fmt"

	"github.com/dittonetwork/executor-avs/cmd/operator/internal/adapters/dittoentrypoint"
	"github.com/dittonetwork/executor-avs/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterExecutor(cfg *CommonFlags) {
	dep, err := initDittoEntryPoint(cfg)
	if err != nil {
		log.With(log.Err(err)).Fatal("ditto entrypoint init error")
	}

	if err = dep.RegisterExecutor(context.Background()); err != nil {
		log.With(log.Err(err)).Fatal("register executor error")
	}
}

func UnregisterExecutor(cfg *CommonFlags) {
	dep, err := initDittoEntryPoint(cfg)
	if err != nil {
		log.With(log.Err(err)).Fatal("ditto entrypoint init error")
	}

	if err = dep.UnregisterExecutor(context.Background()); err != nil {
		log.With(log.Err(err)).Fatal("unregister executor error")
	}
}

func ArrangeExecutors(cfg *CommonFlags) {
	dep, err := initDittoEntryPoint(cfg)
	if err != nil {
		log.With(log.Err(err)).Fatal("ditto entrypoint init error")
	}

	ctx := context.Background()
	executorsAmount, err := dep.GetAmountExecutors(ctx)
	if err != nil {
		log.With(log.Err(err)).Fatal("get amount executors error")
	}

	log.With(log.Int64("executors_amount", executorsAmount.Int64())).Info("executors amount")

	if err = dep.ArrangeExecutors(ctx); err != nil {
		log.With(log.Err(err)).Fatal("arrange executor error")
	}
}

func initDittoEntryPoint(cfg *CommonFlags) (*dittoentrypoint.DittoEntryPoint, error) {
	conn, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		return nil, fmt.Errorf("ethereum client dial: %w", err)
	}

	entryPoint, err := dittoentrypoint.New(conn, cfg.ContractAddress, cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("dittoentrypoint client init: %w", err)
	}

	return entryPoint, nil
}
