package cmd

import (
	"fmt"

	"github.com/LucasRufo/go-kindxt/charts"
	"github.com/LucasRufo/go-kindxt/config"
	"github.com/LucasRufo/go-kindxt/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	rootCmd.AddCommand(createClusterCommand)

	for _, helmPackage := range charts.HelmPackages {
		createClusterCommand.PersistentFlags().Bool(helmPackage.ParameterName, false, helmPackage.Description)
	}
}

func deleteKindCluster() error {
	err := util.ExecuteOsCommand("kind", "delete", "cluster")

	if err != nil {
		return err
	}

	fmt.Println("Cluster deleted.")

	return nil
}

func createKindCluster(extraPortBindings []config.ExtraPortMapping) error {
	fileName := config.CreateKindYAMLConfig(extraPortBindings)

	configFlag := fmt.Sprintf("--config=%s", fileName)

	err := util.ExecuteOsCommand("kind", "create", "cluster", configFlag, "--wait=1m")

	if err != nil {
		return err
	}

	return nil
}

var createClusterCommand = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates an empty Kind cluster. Specify flags if you need dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		//there is no problem in trying to delete the cluster every time that we create a new one.
		deleteKindCluster()

		var helmPackagesToInstall []charts.HelmPackage

		cmd.Flags().Visit(func(f *pflag.Flag) {
			for _, helmPackage := range charts.HelmPackages {
				helmPackagesToInstall = append(helmPackagesToInstall, helmPackage)
			}
		})

		var extraPortBindings []config.ExtraPortMapping

		for _, helmPackage := range helmPackagesToInstall {
			extraPortBindings = append(extraPortBindings, helmPackage.Ports)
		}

		createKindCluster(extraPortBindings)

		for _, helmPackage := range helmPackagesToInstall {
			err := helmPackage.Install()

			fmt.Println(err)
		}
	},
}
