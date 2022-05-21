package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kaiji",
		Short:         "Write a RANDOM selection of the input lines to standard output.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.SetHelpCommand(newCommandHelp())

	cmd.AddCommand(newCommandKeno())
	cmd.AddCommand(newCommandRoulette())

	return cmd
}

// Execute command functionality
func Execute() {
	cmd := newRootCmd()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Printf("%v\n\n", err)
		cmd.Usage()
		os.Exit(1)
	}
}
