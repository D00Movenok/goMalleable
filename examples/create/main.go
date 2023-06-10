package main

import (
	"fmt"

	malleable "github.com/D00Movenok/goMalleable"
)

func main() {
	p := malleable.Profile{
		SampleName: "test.profile",
		SleepTime:  1337,
		Jitter:     50,
		HTTPGet: []malleable.HTTPGet{{
			URI: "/test",
			Client: malleable.HTTPGetClient{
				Headers: []malleable.Header{
					{
						Name:  "Test",
						Value: "header",
					},
					{
						Name:  "User-Agent",
						Value: "curl/1.3.3.7",
					},
				},
				Parameters: []malleable.Parameter{{
					Name:  "testparam",
					Value: "testvalue",
				}},
				Metadata: []malleable.Function{
					{
						Func: "base64",
					},
					{
						Func: "parameter",
						Args: []string{"meta"},
					},
				},
			},
			Server: malleable.HTTPServer{
				Headers: []malleable.Header{{
					Name:  "Server",
					Value: "Nginx",
				}},
				Output: []malleable.Function{
					{
						Func: "xor",
					},
					{
						Func: "print",
					},
				},
			},
		}},
	}

	fmt.Println(p.String())
}
