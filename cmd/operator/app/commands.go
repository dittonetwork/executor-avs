package app

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

type CommonFlags struct {
	NodeURL         string `mapstructure:"node_url"`
	ContractAddress string `mapstructure:"contract_address"`
	PrivateKey      string `mapstructure:"private_key"`
}

func setupRootCommand() *cobra.Command {
	cfg := &CommonFlags{}

	var rootCmd = &cobra.Command{
		Use:   "operator",
		Short: "Operator manages blockchain interactions",
		Long:  "Operator is a CLI tool for managing blockchain interactions efficiently.",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			cfg.PrivateKey = os.Getenv("OPERATOR_PRIVATE_KEY")
			if cfg.PrivateKey == "" {
				return errors.New("private key must be provided via environment variable OPERATOR_PRIVATE_KEY")
			}
			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {},
	}

	rootCmd.PersistentFlags().StringVar(&cfg.NodeURL, "node-url", "", "URL of the blockchain node")
	rootCmd.PersistentFlags().StringVar(&cfg.ContractAddress, "contract-addr", "", "Contract address")

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run executor",
		Run: func(cmd *cobra.Command, _ []string) {
			initRunFlags(cmd)
			wg := Run(cfg)
			wg.Wait()
		},
	}

	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Register executor",
		Run: func(_ *cobra.Command, _ []string) {
			RegisterExecutor(cfg)
		},
	}

	unregisterCmd := &cobra.Command{
		Use:   "unregister",
		Short: "Unregister executor",
		Run: func(_ *cobra.Command, _ []string) {
			UnregisterExecutor(cfg)
		},
	}

	arrangeExecutorsCmd := &cobra.Command{
		Use:   "arrange",
		Short: "Arrange executors",
		Run: func(_ *cobra.Command, _ []string) {
			ArrangeExecutors(cfg)
		},
	}

	rootCmd.AddCommand(runCmd, registerCmd, unregisterCmd, arrangeExecutorsCmd)
	return rootCmd
}

func Execute() {
	var rootCmd = setupRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.With(log.Err(err)).Fatal("app run error")
	}
}
