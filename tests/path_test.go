package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestSearchInSpecificDirectory(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "search.go", "internal/search")
	cmd.Dir = "../"
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Error running search in specific directory: %v", err)
	}

	expectedOutput := "search.go"
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Expected file '%s' not found in output. Got: %s", expectedOutput, output)
	}
}
