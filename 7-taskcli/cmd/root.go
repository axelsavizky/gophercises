package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.\n",
}

func init() {
	RootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
