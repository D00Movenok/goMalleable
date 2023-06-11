package malleable_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	malleable "github.com/D00Movenok/goMalleable"
)

func TestMalleable_ParsePrintParse(t *testing.T) {
	profile := "testdata/sample.profile"

	data, err := os.Open(profile)
	if err != nil {
		t.Error("Error reading testdata")
	}

	p1, err := malleable.Parse(data)
	if err != nil {
		t.Error("Error parsing profile (1)")
	}

	p2, err := malleable.Parse(strings.NewReader(p1.String()))
	if err != nil {
		t.Error("Error parsing profile (2)")
	}

	if !reflect.DeepEqual(p1, p2) {
		t.Error("Profiles are not equal")
	}
}
