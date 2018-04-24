// Package kubectl is a wrapper around a kubectl tool
package kubectl

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runKubectl(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("[%s] %s", err.Error(), stderr.String())
	}
	return nil
}

// Apply a configuration change to a resource from a file or stdin
func Apply(args ...string) error {
	args = append([]string{"apply"}, args...)
	return runKubectl(args...)
}

// Create one or more resources from a file or stdin
func Create(args ...string) error {
	args = append([]string{"create"}, args...)
	return runKubectl(args...)
}

// Delete resources either from a file, stdin, or specifying label selectors,
// names, resource selectors, or resources
func Delete(args ...string) error {
	args = append([]string{"delete"}, args...)
	return runKubectl(args...)
}

// Run a specified image on the cluster
func Run(args ...string) error {
	args = append([]string{"run"}, args...)
	return runKubectl(args...)
}
