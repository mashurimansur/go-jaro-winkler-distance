GO Jaro Winkler Distance

Native [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) in Go.

Jaro-Winkler distance calculates the familiarity of two strings on range 0 to 1.

### Example

```go
	package main

	import (
		"log"

		jaro "github.com/mashurimansur/go-jaro-winkler-distance"
	)

	func main() {
        options := jaro.Options{
            CaseSensitive: true
        }
        distance := jaro.Distance("Hello", "Hello", options)
        log.Println(distance)
        // Output: 0.1
    }