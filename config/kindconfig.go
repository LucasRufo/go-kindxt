package config

import (
	"fmt"
	"io/ioutil"

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

func CreateKindYAMLConfig(extraPortBindings []ExtraPortMapping) string {
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

	for _, e := range extraPortBindings {
		kindConfig.Nodes[0].ExtraPortMappings = append(kindConfig.Nodes[0].ExtraPortMappings, e)
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

	return fileName
}
