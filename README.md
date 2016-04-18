jaccard
===
[![Build Status](https://travis-ci.org/Iwark/jaccard.svg?branch=master)](https://travis-ci.org/Iwark/jaccard)
[![Coverage Status](https://coveralls.io/repos/github/Iwark/jaccard/badge.svg?branch=master)](https://coveralls.io/github/Iwark/jaccard?branch=master)
[![GoReport](https://goreportcard.com/badge/Iwark/jaccard)](http://goreportcard.com/report/Iwark/jaccard)
[![GoDoc](https://godoc.org/github.com/Iwark/jaccard?status.svg)](https://godoc.org/github.com/Iwark/jaccard)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Package `jaccard` provides calcuration of jaccard index.  
``CalcurateBySlices`` is about 5 times faster than using [set](github.com/deckarep/golang-set)

## Installation

```
go get github.com/Iwark/jaccard
```

## Example

```go
package main

import (
  "fmt"
  "github.com/Iwark/jaccard"
)

func main() {
  s1 := []interface{}{1, 2, 3, 4}
  s2 := []interface{}{1, 3, 5, 7}

  d := jaccard.CalcurateBySlices(s1, s2)
  fmt.Println("jaccard index: ", d)
}
```

## License

jaccard is released under the MIT License.
