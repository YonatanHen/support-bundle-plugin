package commands

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBodyGenerator(t *testing.T) {
	var requestBody, _ = requestBodyGenerator("../utils/test.txt")
	assert.NotEmpty(t, requestBody)
}

func TestFileUpload(t *testing.T) {
	var requestBody, _ = requestBodyGenerator("../utils/test.txt")
	req, err := http.NewRequest("PUT", "https://supportlogs.jfrog.com/logs/1/", requestBody)
	if err != nil {
		t.Fatalf("Error creating the request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code should be 200")

}
