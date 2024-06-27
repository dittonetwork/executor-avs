package app

import (
	"errors"
	"fmt"
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
	if err := rootCmd.MarkPersistentFlagRequired("node-url"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %s", err)
	}
	rootCmd.PersistentFlags().StringVar(&cfg.ContractAddress, "contract-addr", "", "contract address")
	if err := rootCmd.MarkPersistentFlagRequired("contract-addr"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %s", err)
	}

	// TODO: return errors from CommandFuncs, so the exit code would be non-zero in case of failure
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run executor",
		Run: func(_ *cobra.Command, _ []string) {
			wg := Run(cfg)
			wg.Wait()
		},
	}
	initRunFlags(runCmd)
	rootCmd.AddCommand(runCmd)

	setupAuxCommands(rootCmd, cfg)

	return rootCmd
}

func setupAuxCommands(rootCmd *cobra.Command, cfg *CommonFlags) {
	registerCmd := &cobra.Command{
		Use:          "register",
		Short:        "Register operator to AVS",
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return RegisterOperator(cfg)
		},
	}

	deregisterCmd := &cobra.Command{
		Use:          "deregister",
		Short:        "Deregister operator from AVS",
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return DeregisterOperator(cfg)
		},
	}

	setSignerCmd := &cobra.Command{
		Use:          "set-signer",
		Short:        "Set delegated signer for routine operations",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return fmt.Errorf("error retrieving address arg: %w", err)
			}
			return SetDelegatedSigner(cfg, address)
		},
	}
	setSignerCmd.Flags().String("address", "", "delegated signer address")
	if err := setSignerCmd.MarkFlagRequired("address"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %s", err)
	}

	activateCmd := &cobra.Command{
		Use:          "activate",
		Short:        "Activate executor",
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return ActivateExecutor(cfg)
		},
	}

	deactivateCmd := &cobra.Command{
		Use:          "deactivate",
		Short:        "Deactivate executor",
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return DeactivateExecutor(cfg)
		},
	}

	arrangeExecutorsCmd := &cobra.Command{
		Use:   "arrange",
		Short: "Arrange executors",
		RunE: func(_ *cobra.Command, _ []string) error {
			return ArrangeExecutors(cfg)
		},
	}

	rootCmd.AddCommand(
		registerCmd,
		deregisterCmd,
		activateCmd,
		deactivateCmd,
		setSignerCmd,
		arrangeExecutorsCmd,
	)
}

func Execute() {
	var rootCmd = setupRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.With(log.Err(err)).Fatal("app run error")
	}
}
