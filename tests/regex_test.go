package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestRegexSearch(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "--regex", ".*\\.md$")
	cmd.Dir = "../"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running regex search: %v", err)
	}

	expectedOutputs := []string{"README.md", "CHANGELOG.md", "CONTRIBUTING.md"}
	for _, eo := range expectedOutputs {
		if !strings.Contains(string(output), eo) {
			t.Errorf("Expected file '%s' not found in regex search output. Got: %s", eo, output)
		}
	}
}
