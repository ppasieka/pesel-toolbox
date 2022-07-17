[![Go](https://github.com/ppasieka/pesel-toolbox/actions/workflows/go.yml/badge.svg)](https://github.com/ppasieka/pesel-toolbox/actions/workflows/go.yml)

# PESEL toolbox

The PESEL (Personal Identification Number) toolbox is a Golang library that helps you with validation and generation of PESEL strings.

It can help you to:

- Validate PESEL strings

- Generate PESEL numbers (Not implemented yet)

## Validate PESEL

Example of PESEL validation

```go
package main

import (
	"fmt"
	"os"

	"github.com/ppasieka/pesel-toolbox"
)

func main() {
	number := os.Args[1]

	pesel, err := pesel.New(number)
	if err != nil {
		fmt.Println(err.Error()) 
	} else {
		fmt.Printf("'%s' is a valid PESEL.", number) 
    }
}
```
