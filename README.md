# goMalleable

🔎🪲 Malleable C2 profiles parser and assembler written in golang

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Examples](#examples)

## Introduction

goMalleable package aims to parse and assemble new CobaltStrike's Malleable C2 profiles.

## Installation

Package can be installed with:

```bash
go get github.com/D00Movenok/goMalleable
```

## Usage

Function `Parse` parses Malleable profile string to easy-to-read structure

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

Example of structure can be found [here](https://github.com/D00Movenok/goMalleable/blob/main/examples/create/main.go#L11), definition of structure can be found [here](https://github.com/D00Movenok/goMalleable/blob/main/readable.go#L9).

You also may print this structure as string to get Malleable profile file. Example: [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/create)

## Examples

| Link | Description |
| ---- | ----------- |
| [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/parse) | Example of profile parsing |
| [Link](https://github.com/D00Movenok/goMalleable/tree/main/examples/create) | Example of profile creation |
