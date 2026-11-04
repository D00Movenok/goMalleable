package malleable

import (
	"fmt"
	"strconv"
	"strings"
)

// Boolean is a custom bool type, because github.com/alecthomas/participle/v2
// parses default bool type as true if something is found.
type Boolean bool

func (b *Boolean) Capture(values []string) error {
	out, err := strconv.ParseBool(values[0])
	if err != nil {
		return err
	}

	*b = (Boolean)(out)

	return nil
}

// CommaSeparatedList is a default comma-separated string list parser and
// stringer, e.g. curl*,lynx*,wget*.
type CommaSeparatedList []string //nolint:recvcheck // mixed receivers: Capture is pointer, String is value

func (l *CommaSeparatedList) Capture(values []string) error {
	s := strings.Split(values[0], ",")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	*l = s
	return nil
}

func (l CommaSeparatedList) String() string {
	return strings.Join(([]string)(l), ",")
}

// SpaceSeparatedList is a default space-separated string list parser and
// stringer, e.g. /jquery-3.3.1.min.js /jquery-1.3.3.7.min.js /someotherurl.
type SpaceSeparatedList []string //nolint:recvcheck // mixed receivers: Capture is pointer, String is value

func (l *SpaceSeparatedList) Capture(values []string) error {
	s := strings.Split(values[0], " ")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	*l = s
	return nil
}

func (l SpaceSeparatedList) String() string {
	return strings.Join(([]string)(l), " ")
}

// Header is a key-value type with "header" prefix, used for headers parsing
// and (mostly) stringer, e.g. header "Accept-Encoding" "gzip, deflate";.
type Header struct {
	Name  string `parser:"@String"`
	Value string `parser:"@String"`
}

func (h Header) String() string {
	return fmt.Sprintf("header %s %s;\n", strconv.Quote(h.Name), strconv.Quote(h.Value))
}

// Parameter is a key-value type with "parameter" prefix, used for parameters
// parsing and (mostly) stringer, e.g. parameter "param_name" "param_value";.
type Parameter struct {
	Name  string `parser:"@String"`
	Value string `parser:"@String"`
}

func (p Parameter) String() string {
	return fmt.Sprintf("parameter %s %s;\n", strconv.Quote(p.Name), strconv.Quote(p.Value))
}

// String is a parser and stringer for the "string" function.
type String string

func (s String) String() string {
	return fmt.Sprintf("string %v;\n", strconv.Quote((string)(s)))
}

// StringW is a parser and stringer for the "stringw" function.
type StringW string

func (s StringW) String() string {
	return fmt.Sprintf("stringw %v;\n", strconv.Quote((string)(s)))
}

// Data is a parser and stringer for the "data" function.
type Data string

func (s Data) String() string {
	return fmt.Sprintf("data %v;\n", strconv.Quote((string)(s)))
}

// Function is a parser and stringer for function sequences, e.g.
// http-get output, transforms in post-ex, etc.
type Function struct {
	Func string   `parser:"@Ident"`
	Args []string `parser:"@String* \";\""`
}

func (f Function) String() string {
	var p strings.Builder
	p.WriteString(f.Func)
	for _, n := range f.Args {
		fmt.Fprintf(&p, " %s", strconv.Quote(n))
	}
	p.WriteString(";\n")
	return p.String()
}
