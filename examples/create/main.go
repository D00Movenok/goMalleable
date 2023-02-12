package main

import (
	"fmt"

	parser "github.com/D00Movenok/goMalleable"
)

func main() {
	// profile from example/parse/test.profile
	profile := &parser.Profile{
		Globals: map[string]string{
			"data_jitter":      "50",
			"host_stage":       "false",
			"jitter":           "33",
			"pipename":         "ntsvcs##",
			"pipename_stager":  "scerpc##",
			"sample_name":      "whatever.profile",
			"sleeptime":        "37500",
			"smb_frame_header": "",
			"ssh_banner":       "Welcome to Ubuntu 18.04.4 LTS (GNU/Linux 4.15.0-1065-aws x86_64)",
			"ssh_pipename":     "SearchTextHarvester##",
			"tcp_frame_header": "",
			"tcp_port":         "8000",
			"useragent":        "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/587.38 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
		},
		HttpsCertificate: map[string]string{
			"C":        "US",
			"CN":       "whatever.com",
			"L":        "California",
			"O":        "whatever LLC.",
			"OU":       "local.org",
			"ST":       "CA",
			"validity": "365",
		},
		CodeSigner: map[string]string{
			"alias":    "server",
			"keystore": "your_keystore.jks",
			"password": "your_password",
		},
		HttpConfig: &parser.HttpConfig{
			Params: map[string]string{
				"block_useragents":      "curl*,lynx*,wget*",
				"headers":               "Server, Content-Type",
				"trust_x_forwarded_for": "false",
			},
			Headers: [][2]string{
				{
					"Server",
					"nginx",
				},
			},
		},
		DnsBeacon: map[string]map[string]string{
			"default": {
				"beacon":             "d-bx.",
				"dns_idle":           "8.8.8.8",
				"dns_max_txt":        "220",
				"dns_sleep":          "0",
				"dns_stager_prepend": ".wwwds.",
				"dns_stager_subhost": ".e2867.dsca.",
				"dns_ttl":            "1",
				"get_A":              "d-1ax.",
				"get_AAAA":           "d-4ax.",
				"get_TXT":            "d-1tx.",
				"maxdns":             "255",
				"ns_response":        "zero",
				"put_metadata":       "d-1mx",
				"put_output":         "d-1ox.",
			},
		},
		HttpGet: map[string]*parser.HttpGet{
			"default": {
				Params: map[string]string{
					"uri": "/login /config /admin",
				},
				Client: &parser.HttpGetClient{
					Headers: [][2]string{
						{
							"Host",
							"whatever.com",
						},
						{
							"Connection",
							"close",
						},
					},
					URIParams: [][2]string{
						{
							"test1",
							"test2",
						},
					},
					Metadata: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "append",
							Values: []string{
								".php",
							},
						},
						{
							Verb: "parameter",
							Values: []string{
								"file",
							},
						},
					},
				},
				Server: &parser.HttpServerOutput{
					Output: []parser.MultiParam{
						{
							Verb: "netbios",
						},
						{
							Verb: "prepend",
							Values: []string{
								"content=",
							},
						},
						{
							Verb: "append",
							Values: []string{
								"\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n",
							},
						},
						{
							Verb: "print",
						},
					},
				},
			},
			"variant_name_get": {
				Params: map[string]string{
					"uri": "/uri1 /uri2 /uri3",
				},
				Client: &parser.HttpGetClient{
					Headers: [][2]string{
						{
							"Host",
							"whatever.com",
						},
						{
							"Connection",
							"close",
						},
					},
					URIParams: [][2]string{
						{
							"test1",
							"test2",
						},
					},
					Metadata: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "append",
							Values: []string{
								".php",
							},
						},
						{
							Verb: "parameter",
							Values: []string{
								"file",
							},
						},
					},
				},
				Server: &parser.HttpServerOutput{
					Headers: [][2]string{
						{
							"Server",
							"nginx",
						},
					},
					Output: []parser.MultiParam{
						{
							Verb: "netbios",
						},
						{
							Verb: "prepend",
							Values: []string{
								"content=",
							},
						},
						{
							Verb: "append",
							Values: []string{
								"\n<meta name=\n",
							},
						},
						{
							Verb: "print",
						},
					},
				},
			},
		},
		HttpPost: map[string]*parser.HttpPost{
			"default": {
				Params: map[string]string{
					"uri":  "/Login /Config /Admin",
					"verb": "GET",
				},
				Client: &parser.HttpPostClient{
					Headers: [][2]string{
						{
							"Host",
							"whatever.com",
						},
						{
							"Connection",
							"close",
						},
					},
					URIParams: [][2]string{
						{
							"test1",
							"test2",
						},
					},
					Output: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "parameter",
							Values: []string{
								"testParam",
							},
						},
					},
					ID: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "parameter",
							Values: []string{
								"id",
							},
						},
					},
				},
				Server: &parser.HttpServerOutput{
					Output: []parser.MultiParam{
						{
							Verb: "netbios",
						},
						{
							Verb: "prepend",
							Values: []string{
								"content=",
							},
						},
						{
							Verb: "append",
							Values: []string{
								"\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n",
							},
						},
						{
							Verb: "print",
						},
					},
				},
			},
			"variant_name_post": {
				Params: map[string]string{
					"uri":  "/Uri1 /Uri2 /Uri3",
					"verb": "GET",
				},
				Client: &parser.HttpPostClient{
					Headers: [][2]string{
						{
							"Host",
							"whatever.com",
						},
						{
							"Connection",
							"close",
						},
					},
					URIParams: [][2]string{},
					Output: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "parameter",
							Values: []string{
								"testParam",
							},
						},
					},
					ID: []parser.MultiParam{
						{
							Verb: "base64url",
						},
						{
							Verb: "parameter",
							Values: []string{
								"id",
							},
						},
					},
				},
				Server: &parser.HttpServerOutput{
					Output: []parser.MultiParam{
						{
							Verb: "netbios",
						},
						{
							Verb: "prepend",
							Values: []string{
								"content=",
							},
						},
						{
							Verb: "append",
							Values: []string{
								"\n<meta name=\n",
							},
						},
						{
							Verb: "print",
						},
					},
				},
			},
		},
		HttpStager: map[string]*parser.HttpStager{
			"default": {
				Params: map[string]string{
					"uri_x64": "/console",
					"uri_x86": "/Console",
				},
				Client: &parser.HttpStagerClient{
					Headers: [][2]string{
						{
							"Host",
							"whatever.com",
						},
						{
							"Connection",
							"close",
						},
					},
					URIParams: [][2]string{
						{
							"test1",
							"test2",
						},
					},
				},
				Server: &parser.HttpServerOutput{
					Output: []parser.MultiParam{
						{
							Verb: "prepend",
							Values: []string{
								"content=",
							},
						},
						{
							Verb: "append",
							Values: []string{
								"</script>\n",
							},
						},
						{
							Verb: "print",
						},
					},
				},
			},
		},
		Stage: &parser.Stage{
			Params: map[string]string{
				"checksum":     "0",
				"cleanup":      "true",
				"compile_time": "25 Oct 2016 01:57:23",
				"entry_point":  "170000",
				"module_x64":   "wwanmm.dll",
				"module_x86":   "wwanmm.dll",
				"obfuscate":    "true",
				"rich_header":  "",
				"sleep_mask":   "true",
				"smartinject":  "true",
				"stomppe":      "true",
				"userwx":       "false",
			},
			String: []string{
				"something",
			},
			Data: []string{
				"something",
			},
			Stringw: []string{
				"something",
			},
			TransformX86: []parser.MultiParam{
				{
					Verb: "prepend",
					Values: []string{
						"\u0090\u0090\u0090",
					},
				},
				{
					Verb: "strrep",
					Values: []string{
						"ReflectiveLoader",
						"",
					},
				},
				{
					Verb: "strrep",
					Values: []string{
						"beacon.dll",
						"",
					},
				},
			},
			TransformX64: []parser.MultiParam{
				{
					Verb: "prepend",
					Values: []string{
						"\u0090\u0090\u0090",
					},
				},
				{
					Verb: "strrep",
					Values: []string{
						"ReflectiveLoader",
						"",
					},
				},
				{
					Verb: "strrep",
					Values: []string{
						"beacon.x64.dll",
						"",
					},
				},
			},
		},
		ProcessInject: &parser.ProcessInject{
			Params: map[string]string{
				"allocator": "NtMapViewOfSection",
				"min_alloc": "16700",
				"startrwx":  "true",
				"userwx":    "false",
			},
			TransformX86: []parser.MultiParam{
				{
					Verb: "prepend",
					Values: []string{
						"\u0090\u0090\u0090",
					},
				},
			},
			TransformX64: []parser.MultiParam{
				{
					Verb: "prepend",
					Values: []string{
						"\u0090\u0090\u0090",
					},
				},
			},
			Execute: []parser.MultiParam{
				{
					Verb: "CreateThread",
					Values: []string{
						"ntdll.dll!RtlUserThreadStart+0x1000",
					},
				},
				{
					Verb: "SetThreadContext",
				},
				{
					Verb: "NtQueueApcThread-s",
				},
				{
					Verb: "CreateRemoteThread",
					Values: []string{
						"kernel32.dll!LoadLibraryA+0x1000",
					},
				},
				{
					Verb: "RtlCreateUserThread",
				},
			},
		},
		PostEx: map[string]string{
			"amsi_disable": "true",
			"keylogger":    "SetWindowsHookEx",
			"obfuscate":    "true",
			"pipename":     "DserNamePipe##",
			"smartinject":  "true",
			"spawnto_x64":  "%windir%\\sysnative\\gpupdate.exe",
			"spawnto_x86":  "%windir%\\syswow64\\gpupdate.exe",
			"thread_hint":  "ntdll.dll!RtlUserThreadStart",
		},
	}

	fmt.Println(profile)
}
