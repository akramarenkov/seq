# Seq

[![Go Reference](https://pkg.go.dev/badge/github.com/akramarenkov/seq.svg)](https://pkg.go.dev/github.com/akramarenkov/seq)
[![Go Report Card](https://goreportcard.com/badge/github.com/akramarenkov/seq)](https://goreportcard.com/report/github.com/akramarenkov/seq)
[![codecov](https://codecov.io/gh/akramarenkov/seq/branch/master/graph/badge.svg?token=)](https://codecov.io/gh/akramarenkov/seq)

## Purpose

Library that allows you to generate a sequence of something

## Usage

Example:

```go
package main

import (
    "fmt"

    "github.com/akramarenkov/seq"
)

func main() {
    fmt.Println(seq.Int(1, 8, 3))
    // Output:
    // [1 4 7]
}
```
