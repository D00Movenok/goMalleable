package malleable

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func getTabs(n int) string {
	var t strings.Builder
	for i := 0; i < n; i++ {
		t.WriteString("\t")
	}
	return t.String()
}

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
			if sft[1] != "\"set\"" {
				continue
			}
			u, _ := strconv.Unquote(sft[2])
			out.WriteString(fmt.Sprintf("%sset %s %v;\n", t, u, strconv.Quote(fmt.Sprint(v))))
		case reflect.Slice:
			tt := t
			if sft[2] == "\"{\"" {
				u, _ := strconv.Unquote(sft[1])
				out.WriteString(fmt.Sprintf("\n%s%s {\n", t, u))
				tt += "\t"
			}

			for j := 0; j < v.Len(); j++ {
				out.WriteString(fmt.Sprintf("%s%s", tt, v.Index(j)))
			}

			if sft[2] == "\"{\"" {
				out.WriteString(fmt.Sprintf("%s}\n", t))
			}
		case reflect.Struct:
			out.WriteString(fmt.Sprintf("%s%s", t, v))
		}
	}

	return out.String()
}

func printUnnamed(n int, block string, s any) string {
	return printNamed(n, block, "", s)
}

func printNamed(n int, block string, name string, s any) string {
	t := getTabs(n)
	var out strings.Builder
	out.WriteString(fmt.Sprintf("\n%s%s", t, block))
	if name != "" {
		out.WriteString(fmt.Sprintf(" %s", strconv.Quote(name)))
	}
	out.WriteString(" {\n")
	out.WriteString(printStruct(n+1, s))
	out.WriteString(fmt.Sprintf("%s}\n", t))
	return out.String()
}
