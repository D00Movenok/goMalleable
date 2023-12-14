# goMalleable

[![PkgGoDev](https://pkg.go.dev/badge/github.com/D00Movenok/goMalleable)](https://pkg.go.dev/github.com/D00Movenok/goMalleable)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/D00Movenok/goMalleable)](https://goreportcard.com/report/github.com/D00Movenok/goMalleable)
[![Test](https://github.com/D00Movenok/goMalleable/actions/workflows/test.yml/badge.svg)](https://github.com/D00Movenok/goMalleable/actions/workflows/test.yml)
[![CodeQL](https://github.com/D00Movenok/goMalleable/actions/workflows/codeql.yml/badge.svg)](https://github.com/D00Movenok/goMalleable/actions/workflows/codeql.yml)

🔎🪲 Malleable C2 profiles parser and assembler library written in golang

**Latest supported CobaltStrike version:** 4.9.1

## Table of Contents

1. [WARNING](#warning)
2. [Installation](#installation)
3. [Usage](#usage)
    1. [Parse](#parse)
    2. [Assembly](#assembly)
4. [Examples](#examples)
5. [TODO](#todo)

## WARNING

goMalleable treats you as a consenting adult and assumes you know how to write Malleable C2 Profiles. It's able to detect syntax errors, however there are no runtime checks implemented. It'll gladly generate profiles that don't actually work in production if instructed to do so. Always run the generated profiles through [c2lint](https://www.cobaltstrike.com/help-malleable-c2) before using them in production!

## Installation

Package can be installed with:

```bash
go get github.com/D00Movenok/goMalleable@v1
```

## Usage

### Parse

Function `Parse` parses Malleable profile string to easy-to-read structure. Full example [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/parse).

```go
package main

import (
    "os"
    malleable "github.com/D00Movenok/goMalleable"
)

func main() {
    ...
    data, _ := os.Open("example.profile")
    parsed, _ := malleable.Parse(data)
    ...
}
```

Full definition of structure can be found [here](https://github.com/D00Movenok/goMalleable/blob/main/profile.go).

### Assembly

You may print this structure as string to get Malleable profile file. Full example: [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/create).

```go
fmt.Println(parsed)
```

Output:

```text
...

set host_stage "false";
set jitter "33";
set tcp_frame_header "";
set useragent "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/587.38 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36";

https-certificate {
    set CN "whatever.com";
    set L "California";
    set O "whatever LLC.";
    set OU "local.org";
    set ST "CA";
    set validity "365";
    set C "US";
}

...
```

## Examples

| Link | Description |
| ---- | ----------- |
| [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/parse) | Example of profile parsing |
| [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/create) | Example of profile creation |

## TODO

- [ ] Use map[Name]Type instead of []Type with Name field
