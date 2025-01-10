# Seq

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/seq.svg)](https://pkg.go.dev/github.com/akramarenkov/seq)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/seq)](https://goreportcard.com/report/github.com/akramarenkov/seq)
[![Coverage Status](https://coveralls.io/repos/github/akramarenkov/seq/badge.svg)](https://coveralls.io/github/akramarenkov/seq)

## Purpose

Library that provides to create a sequence of something

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/seq"
)

func main() {
    fmt.Println(seq.Linear(1, 8, 3))
    // Output:
    // [1 4 7 8]
}
```
