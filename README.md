jaccard
===
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
  s1 := []interface{1, 2, 3, 4}
  s2 := []interface{1, 3, 5, 7}

  d := jaccard.CalcurateBySlices(s1, s2)
  fmt.Println("jaccard index: ", d)
}
```

## License

jaccard is released under the MIT License.
