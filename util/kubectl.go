package util

import (
	"bytes"
	"os/exec"
)

func runKubectl(args ...string) (bytes.Buffer, error) {
	cmd := exec.Command("kubectl", args...)
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stderr, err
}

// Apply a configuration change to a resource from a file or stdin
func Apply(filename string, args ...string) (bytes.Buffer, error) {
	args = append([]string{"create", "-f", filename}, args...)
	return runKubectl(args...)
}

// Create one or more resources from a file or stdin
func Create(filename string, args ...string) (bytes.Buffer, error) {
	args = append([]string{"create", "-f", filename}, args...)
	return runKubectl(args...)
}

// Delete resources either from a file, stdin, or specifying label selectors,
// names, resource selectors, or resources
func Delete(filename string, args ...string) (bytes.Buffer, error) {
	args = append([]string{"delete", "-f", filename}, args...)
	return runKubectl(args...)
}

// Run a specified image on the cluster
func Run(args ...string) (bytes.Buffer, error) {
	args = append([]string{"run"}, args...)
	return runKubectl(args...)
}
