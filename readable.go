package parser

import (
	"errors"

	"github.com/alecthomas/repr"
)

type Profile struct {
	// global variables like "set kek \"lol\"" out of other blocks
	Globals map[string]string

	// https-certificate block
	HttpsCertificate map[string]string
	// code-signer block
	CodeSigner map[string]string

	// http-config block
	HttpConfig *HttpConfig

	// map of dns-beacon profiles
	DnsBeacon map[string]map[string]string
	// map of http-get profiles
	HttpGet map[string]*HttpGet
	// map of http-post profiles
	HttpPost map[string]*HttpPost

	// map of http-stager profiles
	HttpStager map[string]*HttpStager

	// stage block
	Stage *Stage

	// process-inject block
	ProcessInject *ProcessInject
	// post-ex block
	PostEx map[string]string
}

// -------------------
// --- HTTP CONFIG ---
// -------------------

// http-config block
type HttpConfig struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string
	// headers array. Example: [["header1", "value1"], ["header2", "value2"]]
	Headers [][2]string
}

// -------------------
// ----- HTTP GET ----
// -------------------

// http-get block
type HttpGet struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string

	// client block
	Client *HttpGetClient
	// server block
	Server *HttpServerOutput
}

// client block in http-get
type HttpGetClient struct {
	// headers array. Example: [["header1", "value1"], ["header2", "value2"]]
	Headers [][2]string
	// uri params array. Example: [["param1", "value1"], ["param2", "value2"]]
	URIParams [][2]string

	// metadata block
	Metadata []MultiParam
}

// -------------------
// ---- HTTP POST ----
// -------------------

// http-post block
type HttpPost struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string

	// client block
	Client *HttpPostClient
	// server block
	Server *HttpServerOutput
}

// client block in http-post
type HttpPostClient struct {
	// headers array. Example: [["header1", "value1"], ["header2", "value2"]]
	Headers [][2]string
	// uri params array. Example: [["param1", "value1"], ["param2", "value2"]]
	URIParams [][2]string

	// output block
	Output []MultiParam
	// id block
	ID []MultiParam
}

// -------------------
// --- HTTP STAGER ---
// -------------------

// http-stager block
type HttpStager struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string

	// client block
	Client *HttpStagerClient
	// server block
	Server *HttpServerOutput
}

// client block in http-stager
type HttpStagerClient struct {
	// headers array. Example: [["header1", "value1"], ["header2", "value2"]]
	Headers [][2]string
	// uri params array. Example: [["param1", "value1"], ["param2", "value2"]]
	URIParams [][2]string
}

// -------------------
// ------ STAGE ------
// -------------------

// stage block
type Stage struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string
	// string params
	String []string
	// data params
	Data []string
	// stringw params
	Stringw []string

	// transform-x86 block
	TransformX86 []MultiParam
	// transform-x64 block
	TransformX64 []MultiParam
}

// -------------------
// -- PROCESS INJECT -
// -------------------

// process-inject block
type ProcessInject struct {
	// params like. Example: "set kek \"lol\""
	Params map[string]string

	// transform-x86 block
	TransformX86 []MultiParam
	// transform-x64 block
	TransformX64 []MultiParam
	// execute block
	Execute []MultiParam
}

// -------------------
// ------ OTHER ------
// -------------------

// basic http server output with headers and output block
type HttpServerOutput struct {
	// headers array. Example: [["header1", "value1"], ["header2", "value2"]]
	Headers [][2]string
	// output block
	Output []MultiParam
}

// Examples: "append \"some data\"", "print", "string \"some string\""
type MultiParam struct {
	Verb   string
	Values []string
}

func getProfileName(name string) string {
	var profileName string
	if name == "" {
		profileName = "default"
	} else {
		profileName = name
	}
	return profileName
}

// parse server output format, e.g.
// server {
// 	header "Server" "nginx";
// 	output {
// 		netbios;
// 		prepend "content=";
// 		append "\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n";
// 		print;
// 	}
// }
func parseServerOutput(group *group) (*HttpServerOutput, error) {
	server := &HttpServerOutput{}
	var err error

	for _, entry := range group.Entries {
		if entry.Func != nil && entry.Func.FuncName == "header" {
			if len(entry.Func.Values) == 2 {
				server.Headers = append(server.Headers,
					[2]string{entry.Func.Values[0], entry.Func.Values[1]})
			} else {
				return nil, errors.New("header: have bad params " + repr.String(entry.Func))
			}
		} else if entry.Group != nil && entry.Group.Type == "output" {
			server.Output, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}
	}

	return server, nil
}

// parse format with only 'set', e.g.
// https-certificate {
//     set C "US";
//     set CN "whatever.com";
//     set L "California";
//     set O "whatever LLC.";
//     set OU "local.org";
//     set ST "CA";
//     set validity "365";
// }
func parseOnlyParams(group *group) (map[string]string, error) {
	parsed := map[string]string{}
	for _, entry := range group.Entries {
		if entry.Param == nil {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}

		parsed[entry.Param.Name] = entry.Param.Value
	}
	return parsed, nil
}

// parse multiparams functions, e.g.
// metadata {
// 	base64url;
// 	append ".php";
// 	parameter "file";
// }
func parseOnlyMultiparam(group *group) ([]MultiParam, error) {
	parsed := []MultiParam{}

	for _, entry := range group.Entries {
		if entry.Func == nil {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}

		parsed = append(parsed,
			MultiParam{entry.Func.FuncName, entry.Func.Values})
	}
	return parsed, nil
}

// parse http-config block
func parseHttpConfig(group *group) (*HttpConfig, error) {
	hc := &HttpConfig{
		Params:  map[string]string{},
		Headers: [][2]string{},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			hc.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Func != nil && entry.Func.FuncName == "header" {
			// "header" params
			if len(entry.Func.Values) != 2 {
				return nil, errors.New("header: have bad params " + repr.String(entry.Func))
			}
			hc.Headers = append(hc.Headers,
				[2]string{entry.Func.Values[0], entry.Func.Values[1]})
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}
	}

	return hc, nil
}

// parse http-get blocks
func parseHttpGet(group *group) (*HttpGet, error) {
	var err error
	hg := &HttpGet{
		Params: map[string]string{},
		Client: &HttpGetClient{
			Headers:   [][2]string{},
			URIParams: [][2]string{},
			Metadata:  []MultiParam{},
		},
		Server: &HttpServerOutput{
			Headers: [][2]string{},
			Output:  []MultiParam{},
		},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			hg.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Group != nil {
			// parse blocks
			switch entry.Group.Type {
			case "client":
				for _, subEntry := range entry.Group.Entries {
					if subEntry.Func != nil && subEntry.Func.FuncName == "header" {
						// "header" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hg.Client.Headers = append(hg.Client.Headers,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else if subEntry.Func != nil && subEntry.Func.FuncName == "parameter" {
						// "parameter" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hg.Client.URIParams = append(hg.Client.URIParams,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else if subEntry.Group != nil && subEntry.Group.Type == "metadata" {
						// "metadata" block
						hg.Client.Metadata, err = parseOnlyMultiparam(subEntry.Group)
						if err != nil {
							return nil, err
						}
					} else {
						return nil, errors.New("unknown entry: " + repr.String(subEntry))
					}
				}

			case "server":
				hg.Server, err = parseServerOutput(entry.Group)
				if err != nil {
					return nil, err
				}

			default:
				return nil, errors.New("unknown block: " + repr.String(entry.Group))
			}
		}
	}

	return hg, nil
}

// parse http-post blocks
func parseHttpPost(group *group) (*HttpPost, error) {
	var err error
	hp := &HttpPost{
		Params: map[string]string{},
		Client: &HttpPostClient{
			Headers:   [][2]string{},
			URIParams: [][2]string{},
			Output:    []MultiParam{},
			ID:        []MultiParam{},
		},
		Server: &HttpServerOutput{
			Headers: [][2]string{},
			Output:  []MultiParam{},
		},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			hp.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Group != nil {
			// parse blocks
			switch entry.Group.Type {
			case "client":
				for _, subEntry := range entry.Group.Entries {
					if subEntry.Func != nil && subEntry.Func.FuncName == "header" {
						// "header" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hp.Client.Headers = append(hp.Client.Headers,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else if subEntry.Func != nil && subEntry.Func.FuncName == "parameter" {
						// "parameter" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hp.Client.URIParams = append(hp.Client.URIParams,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else if subEntry.Group != nil && subEntry.Group.Type == "id" {
						// "id" block
						hp.Client.ID, err = parseOnlyMultiparam(subEntry.Group)
						if err != nil {
							return nil, err
						}
					} else if subEntry.Group != nil && subEntry.Group.Type == "output" {
						// "output" block
						hp.Client.Output, err = parseOnlyMultiparam(subEntry.Group)
						if err != nil {
							return nil, err
						}
					} else {
						return nil, errors.New("unknown entry: " + repr.String(subEntry))
					}
				}

			case "server":
				hp.Server, err = parseServerOutput(entry.Group)
				if err != nil {
					return nil, err
				}

			default:
				return nil, errors.New("unknown block: " + repr.String(entry.Group))
			}
		}
	}

	return hp, nil
}

// parse http-stager blocks
func parseHttpStager(group *group) (*HttpStager, error) {
	var err error
	hs := &HttpStager{
		Params: map[string]string{},
		Client: &HttpStagerClient{
			Headers:   [][2]string{},
			URIParams: [][2]string{},
		},
		Server: &HttpServerOutput{
			Headers: [][2]string{},
			Output:  []MultiParam{},
		},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			hs.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Group != nil {
			// parse blocks
			switch entry.Group.Type {
			case "client":
				for _, subEntry := range entry.Group.Entries {
					if subEntry.Func != nil && subEntry.Func.FuncName == "header" {
						// "header" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hs.Client.Headers = append(hs.Client.Headers,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else if subEntry.Func != nil && subEntry.Func.FuncName == "parameter" {
						// "parameter" params
						if len(subEntry.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(subEntry.Func))
						}
						hs.Client.URIParams = append(hs.Client.URIParams,
							[2]string{subEntry.Func.Values[0], subEntry.Func.Values[1]})
					} else {
						return nil, errors.New("unknown entry: " + repr.String(subEntry))
					}
				}

			case "server":
				hs.Server, err = parseServerOutput(entry.Group)
				if err != nil {
					return nil, err
				}

			default:
				return nil, errors.New("unknown block: " + repr.String(entry.Group))
			}
		}
	}

	return hs, nil
}

// parse "stage" block
func parseStage(group *group) (*Stage, error) {
	var err error
	s := &Stage{
		Params:       map[string]string{},
		String:       []string{},
		Data:         []string{},
		Stringw:      []string{},
		TransformX86: []MultiParam{},
		TransformX64: []MultiParam{},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			s.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Func != nil && entry.Func.FuncName == "string" {
			// "string" params
			s.String = append(s.String, entry.Func.Values[0])
		} else if entry.Func != nil && entry.Func.FuncName == "data" {
			// "data" params
			s.Data = append(s.Data, entry.Func.Values[0])
		} else if entry.Func != nil && entry.Func.FuncName == "stringw" {
			// "stringw" params
			s.Stringw = append(s.Stringw, entry.Func.Values[0])
		} else if entry.Group != nil && entry.Group.Type == "transform-x86" {
			// "transform-x86" block
			s.TransformX86, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else if entry.Group != nil && entry.Group.Type == "transform-x64" {
			// "transform-x64" block
			s.TransformX64, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}
	}

	return s, nil
}

// process-inject block
func parseProcessInject(group *group) (*ProcessInject, error) {
	var err error
	pi := &ProcessInject{
		Params:       map[string]string{},
		TransformX86: []MultiParam{},
		TransformX64: []MultiParam{},
	}

	for _, entry := range group.Entries {
		if entry.Param != nil {
			// "set" params
			pi.Params[entry.Param.Name] = entry.Param.Value
		} else if entry.Group != nil && entry.Group.Type == "transform-x86" {
			// "transform-x86" block
			pi.TransformX86, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else if entry.Group != nil && entry.Group.Type == "transform-x64" {
			// "transform-x64" block
			pi.TransformX64, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else if entry.Group != nil && entry.Group.Type == "execute" {
			// "execute" block
			pi.Execute, err = parseOnlyMultiparam(entry.Group)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}
	}

	return pi, err
}

// parse profile stucture to easy-to-read Profile structure
func parseToReadable(p *profile) (*Profile, error) {
	var err error

	parsed := &Profile{
		Globals:    map[string]string{},
		DnsBeacon:  map[string]map[string]string{},
		HttpConfig: &HttpConfig{},
		HttpGet:    map[string]*HttpGet{},
		HttpPost:   map[string]*HttpPost{},
		HttpStager: map[string]*HttpStager{},
	}

	var profileName string

	for _, entry := range p.Entries {
		if entry.Param != nil {
			// global "set" params
			parsed.Globals[entry.Param.Name] = entry.Param.Value
		} else if entry.Group != nil {
			switch entry.Group.Type {
			case "https-certificate":
				parsed.HttpsCertificate, err = parseOnlyParams(entry.Group)
				if err != nil {
					return nil, err
				}

			case "code-signer":
				parsed.CodeSigner, err = parseOnlyParams(entry.Group)
				if err != nil {
					return nil, err
				}

			case "http-config":
				parsed.HttpConfig, err = parseHttpConfig(entry.Group)
				if err != nil {
					return nil, err
				}

			case "dns-beacon":
				profileName = getProfileName(entry.Group.Name)
				parsed.DnsBeacon[profileName], err = parseOnlyParams(entry.Group)
				if err != nil {
					return nil, err
				}

			case "http-get":
				profileName = getProfileName(entry.Group.Name)
				parsed.HttpGet[profileName], err = parseHttpGet(entry.Group)
				if err != nil {
					return nil, err
				}

			case "http-post":
				profileName = getProfileName(entry.Group.Name)
				parsed.HttpPost[profileName], err = parseHttpPost(entry.Group)
				if err != nil {
					return nil, err
				}

			case "http-stager":
				profileName = getProfileName(entry.Group.Name)
				parsed.HttpStager[profileName], err = parseHttpStager(entry.Group)
				if err != nil {
					return nil, err
				}

			case "stage":
				parsed.Stage, err = parseStage(entry.Group)
				if err != nil {
					return nil, err
				}

			case "process-inject":
				parsed.ProcessInject, err = parseProcessInject(entry.Group)
				if err != nil {
					return nil, err
				}

			case "post-ex":
				parsed.PostEx, err = parseOnlyParams(entry.Group)
				if err != nil {
					return nil, err
				}

			default:
				return nil, errors.New("unknown block: " + repr.String(entry.Group))
			}
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry))
		}
	}

	return parsed, nil
}
