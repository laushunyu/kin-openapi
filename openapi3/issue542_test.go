package openapi3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIssue542(t *testing.T) {
	sl := NewLoader()

	_, err := sl.LoadFromFile("testdata/issue542.yml")
	require.NoError(t, err)
}
