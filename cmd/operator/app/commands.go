package app

import (
	"github.com/spf13/cobra"

	"github.com/dittonetwork/executor-avs/pkg/log"
)

var nodeURL, contractAddress, privateKey string

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
			wg := Run()
			wg.Wait()
		},
	}
	registerCmd = &cobra.Command{
		Use:   "register",
		Short: "Register executor",
		Run: func(_ *cobra.Command, _ []string) {
			RegisterExecutor()
		},
	}
	unregisterCmd = &cobra.Command{
		Use:   "unregister",
		Short: "Unregister executor",
		Run: func(_ *cobra.Command, _ []string) {
			UnregisterExecutor()
		},
	}
)

func InitCommands() {
	initCommonFlags(runCmd, registerCmd, unregisterCmd)
	initRunFlags(runCmd)

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(unregisterCmd)
}

func initCommonFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().StringVar(&nodeURL, "node-url", "", "node URL (using websocket)")
		cmd.Flags().StringVar(&contractAddress, "contract-addr", "", "DittoEntryPoint contract address")
		cmd.Flags().StringVar(&privateKey, "private-key", "", "Private key")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.With(log.Err(err)).Fatal("app run error")
	}
}
