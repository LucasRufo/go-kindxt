package mongodb

import (
	"os"
	"path"

	"github.com/LucasRufo/go-kindxt/config"
	"github.com/LucasRufo/go-kindxt/util"
)

var Port = config.ExtraPortMapping{
	ContainerPort: 30001,
	HostPort:      27017,
	Protocol:      "TCP",
}

var ParameterName = "mongodb"

var Description = "Installs MongoDB into the Cluster."

func Install() error {
	currentDirectory, err := os.Getwd()

	if err != nil {
		return err
	}

	configPath := path.Join(currentDirectory, "charts", "mongodb", "config.yaml")

	err = util.InstallFromRepo("bitnami/mongodb", "bitnami", "https://charts.bitnami.com/bitnami", "mongodb", configPath, "default")

	if err != nil {
		return err
	}

	return nil
}
