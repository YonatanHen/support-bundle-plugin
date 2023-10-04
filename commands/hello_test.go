package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestBodyGenerator(t *testing.T) {
	var requestBody, _ = requestBodyGenerator("../utils/test.txt")
	assert.NotEmpty(t, requestBody)
}
