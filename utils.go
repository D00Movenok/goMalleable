package malleable

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// NOTE: index of the option name token in the parser struct tag.
const parserTagOptionIdx = 2

func getTabs(n int) string {
	var t strings.Builder
	for i := 0; i < n; i++ {
		t.WriteString("\t")
	}
	return t.String()
}

// Print structures internals.
func printStruct(n int, s any) string {
	t := getTabs(n)

	var out strings.Builder

	st := reflect.TypeOf(s)
	sv := reflect.ValueOf(s)
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		v := sv.Field(i)

		sft := strings.Split(f.Tag.Get("parser"), " ")
		switch f.Type.Kind() { //nolint: exhaustive // not needed by design
		case reflect.Int, reflect.String, reflect.Bool:
			// skip e.g. names of named blocks, they are printed
			// by their parents' Stringers
			if sft[1] == "\"set\"" {
				printSet(&out, t, sft, v)
			}
		case reflect.Slice:
			if sft[1] == "\"set\"" {
				printSet(&out, t, sft, v)
			} else {
				printSlice(&out, t, sft, v)
			}
		case reflect.Struct:
			fmt.Fprintf(&out, "%s%s", t, v)
		}
	}

	return out.String()
}

// Print "set" fields, e.g. set sleeptime "1000";.
func printSet(out *strings.Builder, t string, sft []string, v reflect.Value) {
	// skip empty list sets e.g. headers_remove, otherwise
	// they will be printed as empty strings and reparsed
	// as lists with a single empty element
	if v.Kind() == reflect.Slice && v.Len() == 0 {
		return
	}

	u, _ := strconv.Unquote(sft[parserTagOptionIdx])
	fmt.Fprintf(out, "%sset %s %v;\n", t, u, strconv.Quote(fmt.Sprint(v)))
}

// Print unnamed slices fields, e.g. StringW or TransformX64.
func printSlice(out *strings.Builder, t string, sft []string, v reflect.Value) {
	// add "name { ... }" for blocked slices
	blocked := len(sft) > parserTagOptionIdx &&
		sft[parserTagOptionIdx] == "\"{\""
	tt := t
	if blocked {
		u, _ := strconv.Unquote(sft[1])
		fmt.Fprintf(out, "\n%s%s {\n", t, u)
		tt += "\t"
	}

	for j := 0; j < v.Len(); j++ {
		fmt.Fprintf(out, "%s%s", tt, v.Index(j))
	}

	if blocked {
		fmt.Fprintf(out, "%s}\n", t)
	}
}

// Print unnamed structures.
func printUnnamed(n int, block string, s any) string {
	return printNamed(n, block, "", s)
}

// Print named structures.
func printNamed(n int, block string, name string, s any) string {
	t := getTabs(n)
	var out strings.Builder
	fmt.Fprintf(&out, "\n%s%s", t, block)
	if name != "" {
		fmt.Fprintf(&out, " %s", strconv.Quote(name))
	}
	out.WriteString(" {\n")
	out.WriteString(printStruct(n+1, s))
	fmt.Fprintf(&out, "%s}\n", t)
	return out.String()
}
