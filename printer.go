package parser

import (
	"strconv"
	"strings"
)

func paramsToString(m map[string]string, n int) string {
	var s strings.Builder

	for k, v := range m {
		s.WriteString(strings.Repeat("\t", n) + "set ")
		s.WriteString(k)
		s.WriteString(" ")
		s.WriteString(strconv.Quote(v))
		s.WriteString(";\n")
	}

	return s.String()
}

func tParamsToString(m [][2]string, w string, n int) string {
	var s strings.Builder

	for _, v := range m {
		s.WriteString(strings.Repeat("\t", n))
		s.WriteString(w)
		s.WriteString(" ")
		s.WriteString(strconv.Quote(v[0]))
		s.WriteString(" ")
		s.WriteString(strconv.Quote(v[1]))
		s.WriteString(";\n")
	}

	return s.String()
}

func generateBlockName(block string, name string) string {
	var s strings.Builder

	s.WriteString(block)
	s.WriteString(" ")
	if name != "default" {
		s.WriteString(strconv.Quote(name))
		s.WriteString(" ")
	}
	s.WriteString("{\n")

	return s.String()
}

func funcBlockToString(m []MultiParam, name string, n int) string {
	var s strings.Builder

	s.WriteString(strings.Repeat("\t", n))
	s.WriteString(name)
	s.WriteString(" {\n")

	for _, v := range m {
		s.WriteString(strings.Repeat("\t", n+1))
		s.WriteString(v.Verb)
		for _, p := range v.Values {
			s.WriteString(" ")
			s.WriteString(strconv.Quote(p))
		}
		s.WriteString(";\n")
	}

	s.WriteString(strings.Repeat("\t", n))
	s.WriteString("}\n")

	return s.String()
}

func stringToString(m []string, f string, n int) string {
	var s strings.Builder

	for _, v := range m {
		s.WriteString(strings.Repeat("\t", n))
		s.WriteString(f)
		s.WriteString(" ")
		s.WriteString(strconv.Quote(v))
		s.WriteString(";\n")
	}

	return s.String()
}

func serverOutputToString(o *HttpServerOutput) string {
	var s strings.Builder

	s.WriteString(strings.Repeat("\t", 1))
	s.WriteString("server {\n")
	s.WriteString(tParamsToString(o.Headers, "header", 2))

	s.WriteString(funcBlockToString(o.Output, "output", 2))

	s.WriteString(strings.Repeat("\t", 1))
	s.WriteString("}\n")

	return s.String()
}

func (p Profile) String() string {
	var s strings.Builder

	if p.Globals != nil {
		s.WriteString(paramsToString(p.Globals, 0))
		s.WriteString("\n")
	}

	if p.HttpsCertificate != nil {
		s.WriteString("https-certificate {\n")
		s.WriteString(paramsToString(p.HttpsCertificate, 1))
		s.WriteString("}\n\n")
	}

	if p.CodeSigner != nil {
		s.WriteString("code-signer {\n")
		s.WriteString(paramsToString(p.CodeSigner, 1))
		s.WriteString("}\n\n")
	}

	if p.HttpConfig != nil {
		s.WriteString("http-config {\n")
		s.WriteString(paramsToString(p.HttpConfig.Params, 1))
		s.WriteString(tParamsToString(p.HttpConfig.Headers, "header", 1))
		s.WriteString("}\n\n")
	}

	if p.DnsBeacon != nil {
		for k, v := range p.DnsBeacon {
			s.WriteString(generateBlockName("dns-beacon", k))
			s.WriteString(paramsToString(v, 1))
			s.WriteString("}\n\n")
		}
	}

	if p.HttpGet != nil {
		for k, v := range p.HttpGet {
			s.WriteString(generateBlockName("http-get", k))

			s.WriteString(paramsToString(v.Params, 1))
			s.WriteString("\n")

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("client {\n")
			s.WriteString(tParamsToString(v.Client.Headers, "header", 2))
			s.WriteString(tParamsToString(v.Client.URIParams, "parameter", 2))

			s.WriteString(funcBlockToString(v.Client.Metadata, "metadata", 2))

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("}\n\n")

			s.WriteString(serverOutputToString(v.Server))

			s.WriteString("}\n\n")
		}
	}

	if p.HttpPost != nil {
		for k, v := range p.HttpPost {
			s.WriteString(generateBlockName("http-post", k))

			s.WriteString(paramsToString(v.Params, 1))
			s.WriteString("\n")

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("client {\n")
			s.WriteString(tParamsToString(v.Client.Headers, "header", 2))
			s.WriteString(tParamsToString(v.Client.URIParams, "parameter", 2))

			s.WriteString(funcBlockToString(v.Client.Output, "output", 2))

			s.WriteString(funcBlockToString(v.Client.ID, "id", 2))

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("}\n\n")

			s.WriteString(serverOutputToString(v.Server))

			s.WriteString("}\n\n")
		}
	}

	if p.HttpStager != nil {
		for k, v := range p.HttpStager {
			s.WriteString(generateBlockName("http-stager", k))

			s.WriteString(paramsToString(v.Params, 1))
			s.WriteString("\n")

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("client {\n")
			s.WriteString(tParamsToString(v.Client.Headers, "header", 2))
			s.WriteString(tParamsToString(v.Client.URIParams, "parameter", 2))

			s.WriteString(strings.Repeat("\t", 1))
			s.WriteString("}\n\n")

			s.WriteString(serverOutputToString(v.Server))

			s.WriteString("}\n\n")
		}
	}

	if p.Stage != nil {
		s.WriteString("stage {\n")

		s.WriteString(paramsToString(p.Stage.Params, 1))
		s.WriteString("\n")

		s.WriteString(stringToString(p.Stage.String, "string", 1))
		s.WriteString(stringToString(p.Stage.Stringw, "stringw", 1))
		s.WriteString(stringToString(p.Stage.Data, "data", 1))
		s.WriteString("\n")

		s.WriteString(funcBlockToString(p.Stage.TransformX86, "transform-x86", 1))

		s.WriteString(funcBlockToString(p.Stage.TransformX64, "transform-x64", 1))

		s.WriteString("}\n\n")
	}

	if p.ProcessInject != nil {
		s.WriteString("process-inject {\n")

		s.WriteString(paramsToString(p.ProcessInject.Params, 1))
		s.WriteString("\n")

		s.WriteString(funcBlockToString(p.ProcessInject.TransformX86, "transform-x86", 1))

		s.WriteString(funcBlockToString(p.ProcessInject.TransformX64, "transform-x64", 1))

		s.WriteString(funcBlockToString(p.ProcessInject.Execute, "execute", 1))

		s.WriteString("}\n\n")
	}

	if p.PostEx != nil {
		s.WriteString("post-ex {\n")
		s.WriteString(paramsToString(p.PostEx, 1))
		s.WriteString("}\n\n")
	}

	return s.String()
}
