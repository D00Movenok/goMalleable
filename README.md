# goMalleable

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/D00Movenok/goMalleable)](https://goreportcard.com/report/github.com/D00Movenok/goMalleable)
[![Tests](https://github.com/D00Movenok/goMalleable/actions/workflows/tests.yml/badge.svg)](https://github.com/D00Movenok/goMalleable/actions/workflows/tests.yml)
[![CodeQL](https://github.com/D00Movenok/goMalleable/actions/workflows/codeql.yml/badge.svg)](https://github.com/D00Movenok/goMalleable/actions/workflows/codeql.yml)

🔎🪲 Malleable C2 profiles parser and assembler written in golang

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
	1. [Parse](#parse)
	2. [Assembly](#assembly)
4. [Examples](#examples)

## Introduction

goMalleable package aims to parse and assemble new CobaltStrike's Malleable C2 profiles.

## Installation

Package can be installed with:

```bash
go get github.com/D00Movenok/goMalleable
```

## Usage

### Parse

Function `Parse` parses Malleable profile string to easy-to-read structure. Full example [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/parse).

```go
package main

import (
    "os"
    parser "github.com/D00Movenok/goMalleable"
)

func main() {
    ...
    data, _ := os.ReadFile("example.profile")
    parsed, _ := parser.Parse(string(data))
    ...
}
```

Example of full structure can be found [here](https://github.com/D00Movenok/goMalleable/blob/main/examples/create/main.go#L11), definition of structure can be found [here](https://github.com/D00Movenok/goMalleable/blob/main/readable.go#L9).

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
