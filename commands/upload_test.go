package commands

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const host = "https://supportlogs.jfrog.com"

func TestSupportLogsURL(t *testing.T) {
	cmd := exec.Command("curl", host, "-v")

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running curl command: %v", err)
	}

	outputStr := string(output)

	assert.Contains(t, outputStr, "HTTP/1.1 302", "Expected a 302 status code")
}

func TestUpload(t *testing.T) {
	cmd := CreateCmdCommand("../utils/test.txt", "https://supportlogs.jfrog.com/logs/1/")

	var err error

	if err = cmd.Run(); err != nil {
		t.Fatalf("Error running curl command: %v", err)
	}

	assert.Equal(t, nil, err)
}
