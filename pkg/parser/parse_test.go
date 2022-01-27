package parser

import (
	"testing"

	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/require"
)

func TestParseJQuerryProfile(t *testing.T) {
	parsed, err := Parse(exampleProfile)
	require.NoError(t, err)
	t.Log(repr.String(parsed))
	require.Equal(t, parsed, realReadable)
}
