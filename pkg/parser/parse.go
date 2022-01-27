package parser

func Parse(data string) (*Profile, error) {
	p, err := preparse(data)
	if err != nil {
		return nil, err
	}

	parsed, err := preparseToReadable(p)

	return parsed, err
}
