package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type KindConfig struct {
	Kind       string  `yaml:"kind"`
	ApiVersion string  `yaml:"apiVersion"`
	Nodes      []Nodes `yaml:"nodes"`
}

type Nodes struct {
	Role                 string             `yaml:"role"`
	Image                string             `yaml:"image"`
	KubeadmConfigPatches []string           `yaml:"kubeadmConfigPatches"`
	ExtraPortMappings    []ExtraPortMapping `yaml:"extraPortMappings"`
}

type ExtraPortMapping struct {
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string `yaml:"protocol"`
}

func init() {
	rootCmd.AddCommand(createClusterCommand)
}

func deleteKindCluster() {
	kindDeleteClusterCmd := exec.Command("kind", "delete", "cluster")

	kindDeleteClusterCmd.Stdout = os.Stdout
	kindDeleteClusterCmd.Stderr = os.Stderr

	starterr := kindDeleteClusterCmd.Start()
	err := kindDeleteClusterCmd.Wait()

	if err != nil || starterr != nil {
		fmt.Println("Cluster deleted.")
	}
}

func createKindCluster() {
	kubeadmConfigPatches := []string{
		`kind: InitConfiguration
nodeRegistration:
  kubeletExtraArgs:
    node-labels: "ingress-ready=true"`,
	}

	kindConfig := KindConfig{
		Kind:       "Cluster",
		ApiVersion: "kind.x-k8s.io/v1alpha4",
		Nodes: []Nodes{
			{
				Role:                 "control-plane",
				Image:                "kindest/node:v1.28.0",
				KubeadmConfigPatches: kubeadmConfigPatches,
				ExtraPortMappings:    []ExtraPortMapping{},
			},
		},
	}

	yamlData, err := yaml.Marshal(&kindConfig)

	if err != nil {
		fmt.Printf("Error while marshaling kind YAML config. %v", err)
	}

	fileName := "kindConfig.yaml"
	err = ioutil.WriteFile(fileName, yamlData, 0644)

	if err != nil {
		panic("Unable to create kind YAML config file.")
	}

	configFlag := fmt.Sprintf("--config=%s", fileName)

	kindCreateClusterCmd := exec.Command("kind", "create", "cluster", configFlag)

	kindCreateClusterCmd.Stdout = os.Stdout
	kindCreateClusterCmd.Stderr = os.Stderr

	kindCreateClusterCmd.Start()
	kindCreateClusterCmd.Wait()
}

var createClusterCommand = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates an empty Kind cluster. Specify flags if you need dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		//there is no problem in trying to delete the cluster every time that we create a new one.
		deleteKindCluster()

		createKindCluster()
	},
}
