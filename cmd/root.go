package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gokindxt",
	Short: "gokindxt is a clone of the amazing kindxt project: https://github.com/sergioprates/kindxt",
	Long:  `gokindxt is a CLI wrapper of the kind project. It helps you setup a local kubernetes cluster with dependencies in minutes.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
