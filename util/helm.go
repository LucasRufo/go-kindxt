package util

import (
	"fmt"
	"strings"
)

func InstallFromRepo(chartName, repoName, repoUrl, releaseName, configPath, namespace string) error {
	err := ExecuteOsCommand("helm", "repo", "add", repoName, repoUrl)

	if err != nil {
		return err
	}

	err = ExecuteOsCommand("helm", "repo", "update")

	if err != nil {
		return err
	}

	err = ExecuteOsCommand("helm", "uninstall", releaseName, "-n", namespace)

	if err != nil {
		return err
	}

	installCommand := fmt.Sprintf("install %s %s -n %s --wait --debug --timeout=5m -f %s", releaseName, chartName, namespace, configPath)

	err = ExecuteOsCommand("helm", strings.Split(installCommand, " ")...)

	if err != nil {
		return err
	}

	return nil
}
