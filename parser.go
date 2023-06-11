package malleable

import (
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// Parse Cobalt-Strike MalleableC2 profile.
func Parse(data io.Reader) (*Profile, error) {
	newLexer := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "Comment", Pattern: `#[^\n]*\n?`},
		{Name: "Whitespace", Pattern: `\s`},

		{Name: "Punct", Pattern: `[{},;]`},
		{Name: "Ident", Pattern: `[a-zA-Z0-9_\-]+`},

		{Name: "Boolean", Pattern: `"(true|false)"`},
		{Name: "Number", Pattern: `"\d+"`},
		{Name: "String", Pattern: `"(\\"|[^"])*"`},
	})

	parser := participle.MustBuild[Profile](
		participle.Lexer(newLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.Unquote("String", "Boolean", "Number"),
		participle.UseLookahead(5),
	)

	a, err := parser.Parse("", data)
	return a, err
}
