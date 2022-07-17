package main

import (
	"fmt"
	"os"

	"github.com/ppasieka/pesel-toolbox"
)

func main() {
	number := os.Args[1]

	p, err := pesel.New(number)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("'%s' is a valid PESEL.\n", p.Number())
	}
}
