package util

import (
	"os"
	"os/exec"
)

func ExecuteOsCommand(command string, args ...string) error {
	c := exec.Command(command, args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Start()

	if err != nil {
		return err
	}

	err = c.Wait()

	if err != nil {
		return err
	}

	return nil
}
