package parser

// parse Malleable C2 profile to *Profile{}
func Parse(data string) (*Profile, error) {
	p, err := preparse(data)
	if err != nil {
		return nil, err
	}

	parsed, err := parseToReadable(p)

	return parsed, err
}
