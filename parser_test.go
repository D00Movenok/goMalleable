package malleable_test

import (
	"os"
	"strings"
	"testing"

	malleable "github.com/D00Movenok/goMalleable"
	"github.com/stretchr/testify/require"
)

func TestMalleable_ParsePrintParse(t *testing.T) {
	profile := "testdata/sample.profile"

	data, err := os.Open(profile)
	require.NoError(t, err, "Error reading testdata")

	p1, err := malleable.Parse(data)
	require.NoError(t, err, "Error parsing profile (1)")

	p2, err := malleable.Parse(strings.NewReader(p1.String()))
	require.NoError(t, err, "Error parsing profile (2)")

	require.Equal(t, p1, p2, "Profiles are not equal")
}
