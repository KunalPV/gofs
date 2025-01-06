package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelpCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--help")

	cmd.Dir = "../"

	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Error running help command: %v", err)
	}

	if !strings.Contains(string(output), "Usage: gofs [options] <pattern> [pathname]") {
		t.Errorf("Expected help output not found. Got: %s", output)
	}
}

func TestVersionCommand(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running version command: %v", err)
	}

	expectedVersion := "gofs version 1.0.0"
	if !strings.Contains(string(output), expectedVersion) {
		t.Errorf("Expected version output '%s' not found. Got: %s", expectedVersion, output)
	}
}
