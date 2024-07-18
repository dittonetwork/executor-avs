package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

type Config struct {
	NodeURL            string
	ContractAddress    string
	OperatorPrivateKey string
	ExecutorPrivateKey string
}

func populateConfigWithEnvs(config *Config) {
	config.OperatorPrivateKey, _ = os.LookupEnv("OPERATOR_PRIVATE_KEY")
	config.ExecutorPrivateKey, _ = os.LookupEnv("EXECUTOR_PRIVATE_KEY")
}

func setupRootCommand() *cobra.Command {
	cfg := &Config{}

	var rootCmd = &cobra.Command{
		Use:   "operator",
		Short: "Ditto AVS Operator",
		// Run:   func(_ *cobra.Command, _ []string) {},
	}
	rootCmd.PersistentFlags().StringVar(&cfg.NodeURL, "node-url", "", "URL of the blockchain node")
	rootCmd.PersistentFlags().StringVar(&cfg.ContractAddress, "contract-addr", "", "contract address")

	populateConfigWithEnvs(cfg)

	// TODO: return errors from CommandFuncs, so the exit code would be non-zero in case of failure
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run executor",
		RunE: func(_ *cobra.Command, _ []string) error {
			wg, err := Run(cfg)
			if err != nil {
				return err
			}
			wg.Wait()
			return nil
		},
	}
	initRunFlags(runCmd)
	rootCmd.AddCommand(runCmd)

	setupAuxCommands(rootCmd, cfg)

	return rootCmd
}

func setupAuxCommands(rootCmd *cobra.Command, cfg *Config) {
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

	generateKey := &cobra.Command{
		Use:   "generate",
		Short: "Generate key pair",
		RunE: func(_ *cobra.Command, _ []string) error {
			return GenerateKey()
		},
	}

	rootCmd.AddCommand(
		registerCmd,
		deregisterCmd,
		activateCmd,
		deactivateCmd,
		setSignerCmd,
		arrangeExecutorsCmd,
		generateKey,
	)
}

func Execute() {
	var rootCmd = setupRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.With(log.Err(err)).Fatal("app run error")
	}
}
