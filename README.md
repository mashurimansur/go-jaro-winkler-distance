GO Jaro Winkler Distance

Native [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) in Go.

Jaro-Winkler distance calculates the familiarity of two strings on range 0 to 1.

### How to install
```bash
go get github.com/mashurimansur/go-jaro-winkler-distance
```

### Example

```go
package main

import (
	"log"

	"github.com/mashurimansur/go-jaro-winkler-distance"
)

func main() {
	options := jaro.Options{
		CaseSensitive: false,
	}

	distance := jaro.Distance("Hello", "Hello", options)
	log.Println(distance) // 1
}

```