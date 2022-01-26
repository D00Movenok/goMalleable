package parser

import (
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type profile struct {
	Entries []*entry `@@+`
}

type group struct {
	Type    string   `@Ident`
	Name    string   `@String?`
	Entries []*entry `"{" @@+ "}"`
}

type entry struct {
	Param *setVar   `  @@`
	Func  *function `| @@`
	Group *group    `| @@`
}

type function struct {
	FuncName string `@Ident`
	Values   string `( @String* ";" )`
}

type setVar struct {
	Name  string `"set" @Ident`
	Value string `( @String ";" )`
}

var (
	newLexer = lexer.MustSimple([]lexer.Rule{
		{"Comment", `#[^\n]*\n?`, nil},
		{"Whitespace", `[ \t\r\n]+`, nil},

		{"Punct", `[{},;]`, nil},

		{`Ident`, `[a-zA-Z0-9_-]+`, nil},
		{`String`, `"(?:\\.|[^"])*"`, nil},
	})

	parser = participle.MustBuild(&profile{},
		participle.Lexer(newLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.Unquote("String"),
		participle.UseLookahead(5),
	)
)

func Parse(data string) (*profile, error) {
	p := &profile{}
	err := parser.Parse("", strings.NewReader(data), p)
	return p, err
}
