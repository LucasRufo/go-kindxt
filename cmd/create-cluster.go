package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createClusterCommand)
}

var createClusterCommand = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates an empty Kind cluster. Specify flags if you need dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Kind cluster")
	},
}
