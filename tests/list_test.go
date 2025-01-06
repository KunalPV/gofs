package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestTopLevelListing(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = "../" // Ensure the working directory is the project root
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running top-level listing: %v", err)
	}

	expectedOutputs := []string{"cmd", "docs", "internal", "scripts", "tests"}
	for _, eo := range expectedOutputs {
		if !strings.Contains(string(output), eo) {
			t.Errorf("Expected directory '%s' not found in output. Got: %s", eo, output)
		}
	}
}
