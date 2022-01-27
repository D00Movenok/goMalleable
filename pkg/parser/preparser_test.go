package parser

import (
	"testing"

	"github.com/alecthomas/repr"
	"github.com/stretchr/testify/require"
)

var realParsed = &profile{
	Entries: []*entry{
		{
			Param: &setVar{
				Name:  "sample_name",
				Value: "whatever.profile",
			},
		},
		{
			Param: &setVar{
				Name:  "sleeptime",
				Value: "37500",
			},
		},
		{
			Param: &setVar{
				Name:  "jitter",
				Value: "33",
			},
		},
		{
			Param: &setVar{
				Name:  "useragent",
				Value: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/587.38 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36",
			},
		},
		{
			Param: &setVar{
				Name:  "data_jitter",
				Value: "50",
			},
		},
		{
			Param: &setVar{
				Name:  "host_stage",
				Value: "false",
			},
		},
		{
			Group: &group{
				Type: "dns-beacon",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "dns_idle",
							Value: "8.8.8.8",
						},
					},
					{
						Param: &setVar{
							Name:  "dns_max_txt",
							Value: "220",
						},
					},
					{
						Param: &setVar{
							Name:  "dns_sleep",
							Value: "0",
						},
					},
					{
						Param: &setVar{
							Name:  "dns_ttl",
							Value: "1",
						},
					},
					{
						Param: &setVar{
							Name:  "maxdns",
							Value: "255",
						},
					},
					{
						Param: &setVar{
							Name:  "dns_stager_prepend",
							Value: ".wwwds.",
						},
					},
					{
						Param: &setVar{
							Name:  "dns_stager_subhost",
							Value: ".e2867.dsca.",
						},
					},
					{
						Param: &setVar{
							Name:  "beacon",
							Value: "d-bx.",
						},
					},
					{
						Param: &setVar{
							Name:  "get_A",
							Value: "d-1ax.",
						},
					},
					{
						Param: &setVar{
							Name:  "get_AAAA",
							Value: "d-4ax.",
						},
					},
					{
						Param: &setVar{
							Name:  "get_TXT",
							Value: "d-1tx.",
						},
					},
					{
						Param: &setVar{
							Name:  "put_metadata",
							Value: "d-1mx",
						},
					},
					{
						Param: &setVar{
							Name:  "put_output",
							Value: "d-1ox.",
						},
					},
					{
						Param: &setVar{
							Name:  "ns_response",
							Value: "zero",
						},
					},
				},
			},
		},
		{
			Param: &setVar{
				Name:  "pipename",
				Value: "ntsvcs##",
			},
		},
		{
			Param: &setVar{
				Name:  "pipename_stager",
				Value: "scerpc##",
			},
		},
		{
			Param: &setVar{
				Name: "smb_frame_header",
			},
		},
		{
			Param: &setVar{
				Name:  "tcp_port",
				Value: "8000",
			},
		},
		{
			Param: &setVar{
				Name: "tcp_frame_header",
			},
		},
		{
			Param: &setVar{
				Name:  "ssh_banner",
				Value: "Welcome to Ubuntu 18.04.4 LTS (GNU/Linux 4.15.0-1065-aws x86_64)",
			},
		},
		{
			Param: &setVar{
				Name:  "ssh_pipename",
				Value: "SearchTextHarvester##",
			},
		},
		{
			Group: &group{
				Type: "https-certificate",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "C",
							Value: "US",
						},
					},
					{
						Param: &setVar{
							Name:  "CN",
							Value: "whatever.com",
						},
					},
					{
						Param: &setVar{
							Name:  "L",
							Value: "California",
						},
					},
					{
						Param: &setVar{
							Name:  "O",
							Value: "whatever LLC.",
						},
					},
					{
						Param: &setVar{
							Name:  "OU",
							Value: "local.org",
						},
					},
					{
						Param: &setVar{
							Name:  "ST",
							Value: "CA",
						},
					},
					{
						Param: &setVar{
							Name:  "validity",
							Value: "365",
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "code-signer",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "keystore",
							Value: "your_keystore.jks",
						},
					},
					{
						Param: &setVar{
							Name:  "password",
							Value: "your_password",
						},
					},
					{
						Param: &setVar{
							Name:  "alias",
							Value: "server",
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-config",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "headers",
							Value: "Server, Content-Type",
						},
					},
					{
						Func: &function{
							FuncName: "header",
							Values: []string{
								"Server",
								"nginx",
							},
						},
					},
					{
						Param: &setVar{
							Name:  "trust_x_forwarded_for",
							Value: "false",
						},
					},
					{
						Param: &setVar{
							Name:  "block_useragents",
							Value: "curl*,lynx*,wget*",
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-get",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "uri",
							Value: "/login /config /admin",
						},
					},
					{
						Group: &group{
							Type: "client",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Host",
											"whatever.com",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Connection",
											"close",
										},
									},
								},
								{
									Group: &group{
										Type: "metadata",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														".php",
													},
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"file",
													},
												},
											},
										},
									},
								},
								{
									Func: &function{
										FuncName: "parameter",
										Values: []string{
											"test1",
											"test2",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "server",
							Entries: []*entry{
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "netbios",
												},
											},
											{
												Func: &function{
													FuncName: "prepend",
													Values: []string{
														"content=",
													},
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														"\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n",
													},
												},
											},
											{
												Func: &function{
													FuncName: "print",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-get",
				Name: "variant_name_get",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "uri",
							Value: "/uri1 /uri2 /uri3",
						},
					},
					{
						Group: &group{
							Type: "client",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Host",
											"whatever.com",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Connection",
											"close",
										},
									},
								},
								{
									Group: &group{
										Type: "metadata",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														".php",
													},
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"file",
													},
												},
											},
										},
									},
								},
								{
									Func: &function{
										FuncName: "parameter",
										Values: []string{
											"test1",
											"test2",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "server",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Server",
											"nginx",
										},
									},
								},
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "netbios",
												},
											},
											{
												Func: &function{
													FuncName: "prepend",
													Values: []string{
														"content=",
													},
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														"\n<meta name=\n",
													},
												},
											},
											{
												Func: &function{
													FuncName: "print",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-post",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "uri",
							Value: "/Login /Config /Admin",
						},
					},
					{
						Param: &setVar{
							Name:  "verb",
							Value: "GET",
						},
					},
					{
						Group: &group{
							Type: "client",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "parameter",
										Values: []string{
											"test1",
											"test2",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Host",
											"whatever.com",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Connection",
											"close",
										},
									},
								},
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"testParam",
													},
												},
											},
										},
									},
								},
								{
									Group: &group{
										Type: "id",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"id",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "server",
							Entries: []*entry{
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "netbios",
												},
											},
											{
												Func: &function{
													FuncName: "prepend",
													Values: []string{
														"content=",
													},
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														"\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n",
													},
												},
											},
											{
												Func: &function{
													FuncName: "print",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-post",
				Name: "variant_name_post",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "uri",
							Value: "/Uri1 /Uri2 /Uri3",
						},
					},
					{
						Param: &setVar{
							Name:  "verb",
							Value: "GET",
						},
					},
					{
						Group: &group{
							Type: "client",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Host",
											"whatever.com",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Connection",
											"close",
										},
									},
								},
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"testParam",
													},
												},
											},
										},
									},
								},
								{
									Group: &group{
										Type: "id",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "base64url",
												},
											},
											{
												Func: &function{
													FuncName: "parameter",
													Values: []string{
														"id",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "server",
							Entries: []*entry{
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "netbios",
												},
											},
											{
												Func: &function{
													FuncName: "prepend",
													Values: []string{
														"content=",
													},
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														"\n<meta name=\n",
													},
												},
											},
											{
												Func: &function{
													FuncName: "print",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "http-stager",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "uri_x86",
							Value: "/Console",
						},
					},
					{
						Param: &setVar{
							Name:  "uri_x64",
							Value: "/console",
						},
					},
					{
						Group: &group{
							Type: "client",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Host",
											"whatever.com",
										},
									},
								},
								{
									Func: &function{
										FuncName: "header",
										Values: []string{
											"Connection",
											"close",
										},
									},
								},
								{
									Func: &function{
										FuncName: "parameter",
										Values: []string{
											"test1",
											"test2",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "server",
							Entries: []*entry{
								{
									Group: &group{
										Type: "output",
										Entries: []*entry{
											{
												Func: &function{
													FuncName: "prepend",
													Values: []string{
														"content=",
													},
												},
											},
											{
												Func: &function{
													FuncName: "append",
													Values: []string{
														"</script>\n",
													},
												},
											},
											{
												Func: &function{
													FuncName: "print",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "stage",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "checksum",
							Value: "0",
						},
					},
					{
						Param: &setVar{
							Name:  "compile_time",
							Value: "25 Oct 2016 01:57:23",
						},
					},
					{
						Param: &setVar{
							Name:  "entry_point",
							Value: "170000",
						},
					},
					{
						Param: &setVar{
							Name:  "userwx",
							Value: "false",
						},
					},
					{
						Param: &setVar{
							Name:  "cleanup",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "sleep_mask",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "stomppe",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "obfuscate",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name: "rich_header",
						},
					},
					{
						Param: &setVar{
							Name:  "sleep_mask",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "smartinject",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "module_x86",
							Value: "wwanmm.dll",
						},
					},
					{
						Param: &setVar{
							Name:  "module_x64",
							Value: "wwanmm.dll",
						},
					},
					{
						Group: &group{
							Type: "transform-x86",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "prepend",
										Values: []string{
											"\u0090\u0090\u0090",
										},
									},
								},
								{
									Func: &function{
										FuncName: "strrep",
										Values: []string{
											"ReflectiveLoader",
											"",
										},
									},
								},
								{
									Func: &function{
										FuncName: "strrep",
										Values: []string{
											"beacon.dll",
											"",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "transform-x64",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "prepend",
										Values: []string{
											"\u0090\u0090\u0090",
										},
									},
								},
								{
									Func: &function{
										FuncName: "strrep",
										Values: []string{
											"ReflectiveLoader",
											"",
										},
									},
								},
								{
									Func: &function{
										FuncName: "strrep",
										Values: []string{
											"beacon.x64.dll",
											"",
										},
									},
								},
							},
						},
					},
					{
						Func: &function{
							FuncName: "string",
							Values: []string{
								"something",
							},
						},
					},
					{
						Func: &function{
							FuncName: "data",
							Values: []string{
								"something",
							},
						},
					},
					{
						Func: &function{
							FuncName: "stringw",
							Values: []string{
								"something",
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "process-inject",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "allocator",
							Value: "NtMapViewOfSection",
						},
					},
					{
						Param: &setVar{
							Name:  "min_alloc",
							Value: "16700",
						},
					},
					{
						Param: &setVar{
							Name:  "userwx",
							Value: "false",
						},
					},
					{
						Param: &setVar{
							Name:  "startrwx",
							Value: "true",
						},
					},
					{
						Group: &group{
							Type: "transform-x86",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "prepend",
										Values: []string{
											"\u0090\u0090\u0090",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "transform-x64",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "prepend",
										Values: []string{
											"\u0090\u0090\u0090",
										},
									},
								},
							},
						},
					},
					{
						Group: &group{
							Type: "execute",
							Entries: []*entry{
								{
									Func: &function{
										FuncName: "CreateThread",
										Values: []string{
											"ntdll.dll!RtlUserThreadStart+0x1000",
										},
									},
								},
								{
									Func: &function{
										FuncName: "SetThreadContext",
									},
								},
								{
									Func: &function{
										FuncName: "NtQueueApcThread-s",
									},
								},
								{
									Func: &function{
										FuncName: "CreateRemoteThread",
										Values: []string{
											"kernel32.dll!LoadLibraryA+0x1000",
										},
									},
								},
								{
									Func: &function{
										FuncName: "RtlCreateUserThread",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Group: &group{
				Type: "post-ex",
				Entries: []*entry{
					{
						Param: &setVar{
							Name:  "spawnto_x86",
							Value: "%windir%\\syswow64\\gpupdate.exe",
						},
					},
					{
						Param: &setVar{
							Name:  "spawnto_x64",
							Value: "%windir%\\sysnative\\gpupdate.exe",
						},
					},
					{
						Param: &setVar{
							Name:  "obfuscate",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "smartinject",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "amsi_disable",
							Value: "true",
						},
					},
					{
						Param: &setVar{
							Name:  "thread_hint",
							Value: "ntdll.dll!RtlUserThreadStart",
						},
					},
					{
						Param: &setVar{
							Name:  "pipename",
							Value: "DserNamePipe##",
						},
					},
					{
						Param: &setVar{
							Name:  "keylogger",
							Value: "SetWindowsHookEx",
						},
					},
				},
			},
		},
	},
}

var exampleProfile = `
#clean template profile - no comments, cleaned up, hopefully easier to build new profiles off of.
#updated with 4.3 options
#xx0hcd

###Global Options###
set sample_name "whatever.profile";

set sleeptime "37500";
set jitter    "33";
set useragent "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/587.38 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36";
set data_jitter "50";

set host_stage "false";

###DNS options###
dns-beacon {
    # Options moved into 'dns-beacon' group in 4.3:
    set dns_idle             "8.8.8.8";
    set dns_max_txt          "220";
    set dns_sleep            "0";
    set dns_ttl              "1";
    set maxdns               "255";
    set dns_stager_prepend   ".wwwds.";
    set dns_stager_subhost   ".e2867.dsca.";
     
    # DNS subhost override options added in 4.3:
    set beacon               "d-bx.";
    set get_A                "d-1ax.";
    set get_AAAA             "d-4ax.";
    set get_TXT              "d-1tx.";
    set put_metadata         "d-1mx";
    set put_output           "d-1ox.";
    set ns_response          "zero";
}

###SMB options###
set pipename "ntsvcs##";
set pipename_stager "scerpc##";
set smb_frame_header "";

###TCP options###
set tcp_port "8000";
set tcp_frame_header "";

###SSH options###
set ssh_banner "Welcome to Ubuntu 18.04.4 LTS (GNU/Linux 4.15.0-1065-aws x86_64)";
set ssh_pipename "SearchTextHarvester##";

###SSL Options###
#https-certificate {
    #set keystore "your_store_file.store";
    #set password "your_store_pass";
#}

https-certificate {
    set C "US";
    set CN "whatever.com";
    set L "California";
    set O "whatever LLC.";
    set OU "local.org";
    set ST "CA";
    set validity "365";
}

code-signer {
    set keystore "your_keystore.jks";
    set password "your_password";
    set alias "server";
}

###HTTP-Config Block###
http-config {
    set headers "Server, Content-Type";
    header "Server" "nginx";

    set trust_x_forwarded_for "false";
    
    set block_useragents "curl*,lynx*,wget*";
}

#set headers_remove "image/x-xbitmap, image/pjpeg, application/vnd";

###HTTP-GET Block###
http-get {

    set uri "/login /config /admin";

    #set verb "POST";
    
    client {

        header "Host" "whatever.com";
        header "Connection" "close";

	   
        metadata {
        #base64
        base64url;
        #mask;
        #netbios;
        #netbiosu;
        #prepend "TEST123";
        append ".php";

        parameter "file";
        #header "Cookie";
        #uri-append;

        #print;
    }

    parameter "test1" "test2";
    }

    server {
        #header "Server" "nginx";
 
        output {

            netbios;
            #netbiosu;
            #base64;
            #base64url;
            #mask;
            	       
	    prepend "content=";

	    append "\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n";

            print;
        }
    }
}

###HTTP-GET VARIANT###
http-get "variant_name_get" {

    set uri "/uri1 /uri2 /uri3";

    #set verb "POST";
    
    client {

        header "Host" "whatever.com";
        header "Connection" "close";

	   
    metadata {

        base64url;
        append ".php";

        parameter "file";
        #header "Cookie";
        #uri-append;

        #print;
    }

    parameter "test1" "test2";
    }

    server {
        header "Server" "nginx";
 
        output {

            netbios;
            	       
            prepend "content=";

            append "\n<meta name=\n";

            print;
        }
    }
}

###HTTP-Post Block###
http-post {
    
    set uri "/Login /Config /Admin";
    set verb "GET";
    #set verb "POST";

    client {
        parameter "test1" "test2";
        header "Host" "whatever.com";
        header "Connection" "close";     
        
        output {
            base64url; 
	    parameter "testParam";
        }

        id {
	    base64url;
	    parameter "id";
            #header "ID-Header";

        }
    }
    

    server {
        #header "Server" "nginx";

        output {
            netbios;	    
	   
	    prepend "content=";

	    append "\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n";

            print;
        }
    }
}

###HTTP-POST VARIANT###
http-post "variant_name_post" {
    
    set uri "/Uri1 /Uri2 /Uri3";
    set verb "GET";
    #set verb "POST";

    client {

	header "Host" "whatever.com";
	header "Connection" "close";     
        
        output {
            base64url; 
	    parameter "testParam";
        }

        id {
	    base64url;
	    parameter "id";

        }
    }

    server {
        #header "Server" "nginx";

        output {
            netbios;	    
	   
	    prepend "content=";

	    append "\n<meta name=\n";

            print;
        }
    }
}

###HTTP-Stager Block###
http-stager {

    set uri_x86 "/Console";
    set uri_x64 "/console";

    client {
        header "Host" "whatever.com";
        header "Connection" "close";
	
	    parameter "test1" "test2";
    }

    server {
        #header "Server" "nginx";
	
	output {
	
	    prepend "content=";
	    
	    append "</script>\n";
	    print;
	}

    }
}


###Malleable PE/Stage Block###
stage {
    set checksum        "0";
    set compile_time    "25 Oct 2016 01:57:23";
    set entry_point     "170000";
    #set image_size_x86 "6586368";
    #set image_size_x64 "6586368";
    #set name	        "WWanMM.dll";
    set userwx 	        "false";
    set cleanup	        "true";
    set sleep_mask	"true";
    set stomppe	        "true";
    set obfuscate	"true";
    set rich_header     "";
    
    #new 4.2. options   
    #set allocator "HeapAlloc";
    #set magic_mz_x86 "MZRE";
    #set magic_mz_x64 "MZAR";
    #set magic_pe "PE";
    
    set sleep_mask "true";
    set smartinject "true";

    set module_x86 "wwanmm.dll";
    set module_x64 "wwanmm.dll";

    transform-x86 {
        prepend "\x90\x90\x90";
        strrep "ReflectiveLoader" "";
        strrep "beacon.dll" "";
        }

    transform-x64 {
        prepend "\x90\x90\x90";
        strrep "ReflectiveLoader" "";
        strrep "beacon.x64.dll" "";
        }

    string "something";
    data "something";
    stringw "something"; 
}

###Process Inject Block###
process-inject {

    set allocator "NtMapViewOfSection";		

    set min_alloc "16700";

    set userwx "false";  
    
    set startrwx "true";
        
    transform-x86 {
        prepend "\x90\x90\x90";
    }
    transform-x64 {
        prepend "\x90\x90\x90";
    }

    execute {
        #CreateThread;
        #CreateRemoteThread;       

        CreateThread "ntdll.dll!RtlUserThreadStart+0x1000";

        SetThreadContext;

        NtQueueApcThread-s;

        #NtQueueApcThread;

        CreateRemoteThread "kernel32.dll!LoadLibraryA+0x1000";

        RtlCreateUserThread;
    }
}

###Post-Ex Block###
post-ex {

    set spawnto_x86 "%windir%\\syswow64\\gpupdate.exe";
    set spawnto_x64 "%windir%\\sysnative\\gpupdate.exe";

    set obfuscate "true";

    set smartinject "true";

    set amsi_disable "true";
    
    #new 4.2 options
    set thread_hint "ntdll.dll!RtlUserThreadStart";
    set pipename "DserNamePipe##";
    set keylogger "SetWindowsHookEx";

}
`

func TestPreparseJQuerryProfile(t *testing.T) {
	parsed, err := preparse(exampleProfile)
	require.NoError(t, err)
	t.Log(repr.String(parsed))
	require.Equal(t, parsed, realParsed)
}
