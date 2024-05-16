package app

import (
	"github.com/spf13/cobra"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

var (
	nodeURL,
	contractAddress string
)

func InitCommands() {
	// TODO: Adding the same flags to different commands requires copy/paste. Find a way to organize it better.
	runCmd.Flags().StringVar(
		&nodeURL,
		"node-url",
		"wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a/",
		"node URL (wss://host/apikey)",
	)
	runCmd.Flags().StringVar(
		&contractAddress,
		"contract-addr",
		"0x9201cFC00bB9fE0477b51560123660183bf2026A",
		"DittoEntryPoint contract address",
	)

	registerCmd.Flags().StringVar(
		&nodeURL,
		"node-url",
		"wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a/",
		"node URL (wss://host/apikey)",
	)
	registerCmd.Flags().StringVar(
		&contractAddress,
		"contract-addr",
		"0x9201cFC00bB9fE0477b51560123660183bf2026A",
		"DittoEntryPoint contract address",
	)

	unregisterCmd.Flags().StringVar(
		&nodeURL,
		"node-url",
		"wss://silent-tame-seed.ethereum-holesky.quiknode.pro/a09b2aafbc9447b172c9964f3ac40c85edf5fd6a/",
		"node URL (wss://host/apikey)",
	)
	unregisterCmd.Flags().StringVar(
		&contractAddress,
		"contract-addr",
		"0x9201cFC00bB9fE0477b51560123660183bf2026A",
		"DittoEntryPoint contract address",
	)

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(unregisterCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:   "operator",
		Short: "",
		Long:  "",
		Run:   func(_ *cobra.Command, _ []string) {},
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run executor",
		Run: func(_ *cobra.Command, _ []string) {
			wg := Run(nodeURL, contractAddress)
			wg.Wait()
		},
	}
	registerCmd = &cobra.Command{
		Use:   "register",
		Short: " Register executor",
		Run: func(_ *cobra.Command, _ []string) {
			RegisterExecutor(nodeURL, contractAddress)
		},
	}
	unregisterCmd = &cobra.Command{
		Use:   "register",
		Short: " Register executor",
		Run: func(_ *cobra.Command, _ []string) {
			UnregisterExecutor(nodeURL, contractAddress)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.With(log.Err(err)).Fatal("app run error")
	}
}
