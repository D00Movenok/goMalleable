package malleable

import (
	"fmt"
	"strconv"
	"strings"
)

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	out, err := strconv.ParseBool(values[0])
	if err != nil {
		return err
	}

	*b = (Boolean)(out)

	return nil
}

type CommaSeparatedList []string

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

type URIs []string

func (l *URIs) Capture(values []string) error {
	s := strings.Split(values[0], " ")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}

	*l = s
	return nil
}

func (l URIs) String() string {
	return strings.Join(([]string)(l), " ")
}

type Header struct {
	Name  string `parser:"@String"`
	Value string `parser:"@String"`
}

func (h Header) String() string {
	return fmt.Sprintf("header %s %s;\n", strconv.Quote(h.Name), strconv.Quote(h.Value))
}

type Parameter struct {
	Name  string `parser:"@String"`
	Value string `parser:"@String"`
}

func (p Parameter) String() string {
	return fmt.Sprintf("parameter %s %s;\n", strconv.Quote(p.Name), strconv.Quote(p.Value))
}

type String string

func (s String) String() string {
	return fmt.Sprintf("string %v;\n", strconv.Quote((string)(s)))
}

type StringW string

func (s StringW) String() string {
	return fmt.Sprintf("stringw %v;\n", strconv.Quote((string)(s)))
}

type Data string

func (s Data) String() string {
	return fmt.Sprintf("data %v;\n", strconv.Quote((string)(s)))
}

type Function struct {
	Func string   `parser:"@Ident"`
	Args []string `parser:"@String* \";\""`
}

func (f Function) String() string {
	var p strings.Builder
	p.WriteString(f.Func)
	for _, n := range f.Args {
		p.WriteString(fmt.Sprintf(" %s", strconv.Quote(n)))
	}
	p.WriteString(";\n")
	return p.String()
}
