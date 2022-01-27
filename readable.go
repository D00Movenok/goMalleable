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

	for _, entry1 := range p.Entries {
		if entry1.Param != nil {
			parsed.Globals[entry1.Param.Name] = entry1.Param.Value
		} else if entry1.Group != nil {
			switch entry1.Group.Type {
			case "https-certificate":
				parsed.HttpsCertificate, err = parseOnlyParams(entry1.Group)
				if err != nil {
					return nil, err
				}

			case "code-signer":
				parsed.CodeSigner, err = parseOnlyParams(entry1.Group)
				if err != nil {
					return nil, err
				}

			case "http-config":
				parsed.HttpConfig = &HttpConfig{
					Params:  map[string]string{},
					Headers: [][2]string{},
				}

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.HttpConfig.Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Func != nil && entry2.Func.FuncName == "header" {
						if len(entry2.Func.Values) != 2 {
							return nil, errors.New("header: have bad params " + repr.String(entry2.Func))
						}
						parsed.HttpConfig.Headers = append(parsed.HttpConfig.Headers,
							[2]string{entry2.Func.Values[0], entry2.Func.Values[1]})
					} else {
						return nil, errors.New("unknown entry: " + repr.String(entry2))
					}
				}

			case "dns-beacon":
				profileName = getProfileName(entry1.Group.Name)
				parsed.DnsBeacon[profileName], err = parseOnlyParams(entry1.Group)
				if err != nil {
					return nil, err
				}

			case "http-get":
				profileName = getProfileName(entry1.Group.Name)

				parsed.HttpGet[profileName] = &HttpGet{
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

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.HttpGet[profileName].Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Group != nil {
						switch entry2.Group.Type {
						case "client":
							for _, entry3 := range entry2.Group.Entries {
								if entry3.Func != nil && entry3.Func.FuncName == "header" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpGet[profileName].Client.Headers = append(parsed.HttpGet[profileName].Client.Headers,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else if entry3.Func != nil && entry3.Func.FuncName == "parameter" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpGet[profileName].Client.URIParams = append(parsed.HttpGet[profileName].Client.URIParams,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else if entry3.Group != nil && entry3.Group.Type == "metadata" {
									parsed.HttpGet[profileName].Client.Metadata, err = parseOnlyMultiparam(entry3.Group)
									if err != nil {
										return nil, err
									}
								} else {
									return nil, errors.New("unknown entry: " + repr.String(entry3))
								}
							}
						case "server":
							parsed.HttpGet[profileName].Server, err = parseServerOutput(entry2.Group)
							if err != nil {
								return nil, err
							}
						default:
							return nil, errors.New("unknown block: " + repr.String(entry2.Group))
						}
					}
				}
			case "http-post":
				profileName = getProfileName(entry1.Group.Name)

				parsed.HttpPost[profileName] = &HttpPost{
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

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.HttpPost[profileName].Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Group != nil {
						switch entry2.Group.Type {
						case "client":
							for _, entry3 := range entry2.Group.Entries {
								if entry3.Func != nil && entry3.Func.FuncName == "header" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpPost[profileName].Client.Headers = append(parsed.HttpPost[profileName].Client.Headers,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else if entry3.Func != nil && entry3.Func.FuncName == "parameter" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpPost[profileName].Client.URIParams = append(parsed.HttpPost[profileName].Client.URIParams,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else if entry3.Group != nil && entry3.Group.Type == "id" {
									parsed.HttpPost[profileName].Client.ID, err = parseOnlyMultiparam(entry3.Group)
									if err != nil {
										return nil, err
									}
								} else if entry3.Group != nil && entry3.Group.Type == "output" {
									parsed.HttpPost[profileName].Client.Output, err = parseOnlyMultiparam(entry3.Group)
									if err != nil {
										return nil, err
									}
								} else {
									return nil, errors.New("unknown entry: " + repr.String(entry3))
								}
							}
						case "server":
							parsed.HttpPost[profileName].Server, err = parseServerOutput(entry2.Group)
							if err != nil {
								return nil, err
							}
						default:
							return nil, errors.New("unknown block: " + repr.String(entry2.Group))
						}
					}
				}

			case "http-stager":
				profileName = getProfileName(entry1.Group.Name)

				parsed.HttpStager[profileName] = &HttpStager{
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

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.HttpStager[profileName].Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Group != nil {
						switch entry2.Group.Type {
						case "client":
							for _, entry3 := range entry2.Group.Entries {
								if entry3.Func != nil && entry3.Func.FuncName == "header" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpStager[profileName].Client.Headers = append(parsed.HttpStager[profileName].Client.Headers,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else if entry3.Func != nil && entry3.Func.FuncName == "parameter" {
									if len(entry3.Func.Values) != 2 {
										return nil, errors.New("header: have bad params " + repr.String(entry3.Func))
									}
									parsed.HttpStager[profileName].Client.URIParams = append(parsed.HttpStager[profileName].Client.URIParams,
										[2]string{entry3.Func.Values[0], entry3.Func.Values[1]})
								} else {
									return nil, errors.New("unknown entry: " + repr.String(entry3))
								}
							}
						case "server":
							parsed.HttpStager[profileName].Server, err = parseServerOutput(entry2.Group)
							if err != nil {
								return nil, err
							}
						default:
							return nil, errors.New("unknown block: " + repr.String(entry2.Group))
						}
					}
				}

			case "stage":
				parsed.Stage = &Stage{
					Params:       map[string]string{},
					String:       []string{},
					Data:         []string{},
					Stringw:      []string{},
					TransformX86: []MultiParam{},
					TransformX64: []MultiParam{},
				}

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.Stage.Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Func != nil && entry2.Func.FuncName == "string" {
						parsed.Stage.String = append(parsed.Stage.String, entry2.Func.Values[0])
					} else if entry2.Func != nil && entry2.Func.FuncName == "data" {
						parsed.Stage.Data = append(parsed.Stage.Data, entry2.Func.Values[0])
					} else if entry2.Func != nil && entry2.Func.FuncName == "stringw" {
						parsed.Stage.Stringw = append(parsed.Stage.Stringw, entry2.Func.Values[0])
					} else if entry2.Group != nil && entry2.Group.Type == "transform-x86" {
						parsed.Stage.TransformX86, err = parseOnlyMultiparam(entry2.Group)
						if err != nil {
							return nil, err
						}
					} else if entry2.Group != nil && entry2.Group.Type == "transform-x64" {
						parsed.Stage.TransformX64, err = parseOnlyMultiparam(entry2.Group)
						if err != nil {
							return nil, err
						}
					} else {
						return nil, errors.New("unknown entry: " + repr.String(entry2))
					}
				}

			case "process-inject":
				parsed.ProcessInject = &ProcessInject{
					Params:       map[string]string{},
					TransformX86: []MultiParam{},
					TransformX64: []MultiParam{},
				}

				for _, entry2 := range entry1.Group.Entries {
					if entry2.Param != nil {
						parsed.ProcessInject.Params[entry2.Param.Name] = entry2.Param.Value
					} else if entry2.Group != nil && entry2.Group.Type == "transform-x86" {
						parsed.ProcessInject.TransformX86, err = parseOnlyMultiparam(entry2.Group)
						if err != nil {
							return nil, err
						}
					} else if entry2.Group != nil && entry2.Group.Type == "transform-x64" {
						parsed.ProcessInject.TransformX64, err = parseOnlyMultiparam(entry2.Group)
						if err != nil {
							return nil, err
						}
					} else if entry2.Group != nil && entry2.Group.Type == "execute" {
						parsed.ProcessInject.Execute, err = parseOnlyMultiparam(entry2.Group)
						if err != nil {
							return nil, err
						}
					} else {
						return nil, errors.New("unknown entry: " + repr.String(entry2))
					}
				}

			case "post-ex":
				parsed.PostEx, err = parseOnlyParams(entry1.Group)
				if err != nil {
					return nil, err
				}

			default:
				return nil, errors.New("unknown block: " + repr.String(entry1.Group))
			}
		} else {
			return nil, errors.New("unknown entry: " + repr.String(entry1))
		}
	}

	return parsed, nil
}
