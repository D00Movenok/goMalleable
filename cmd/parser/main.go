package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/alecthomas/repr"

	"github.com/D00Movenok/goMalleable/pkg/parser"
)

var cli struct {
	Files []string `arg:"" optional:"" type:"existingfile" help:"Malleable C2 profile files to parse."`
}

func main() {
	ctx := kong.Parse(&cli)
	for _, file := range cli.Files {
		data, err := os.ReadFile(file)
		ctx.FatalIfErrorf(err)

		parsed, err := parser.Parse(string(data))
		ctx.FatalIfErrorf(err)

		repr.Println(parsed)
	}
}
