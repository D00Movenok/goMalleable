package main

import (
	"log"
	"os"

	malleable "github.com/D00Movenok/goMalleable"
	"github.com/alecthomas/repr"
	"github.com/spf13/pflag"
)

var (
	path = pflag.StringP("path", "p", "../testdata/sample.profile", "Path to the Malleable .profile file")
)

func main() {
	pflag.Parse()

	data, err := os.Open(*path)
	if err != nil {
		log.Fatalf("Can't read file: %s", err)
	}

	parsed, err := malleable.Parse(data)
	if err != nil {
		log.Fatalf("Can't parse profile: %s", err)
	}

	repr.Println(parsed)
}
