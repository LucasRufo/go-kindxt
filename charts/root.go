package charts

import (
	"github.com/LucasRufo/go-kindxt/charts/mongodb"
	"github.com/LucasRufo/go-kindxt/config"
)

type HelmPackage struct {
	ParameterName string
	Description   string
	Ports         config.ExtraPortMapping
	Install       func() error
}

var HelmPackages = []HelmPackage{
	{
		ParameterName: mongodb.ParameterName,
		Description:   mongodb.Description,
		Ports:         mongodb.Port,
		Install:       mongodb.Install,
	},
}
