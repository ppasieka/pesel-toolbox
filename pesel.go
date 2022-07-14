package pesel

import "fmt"

type Date struct {
	year  uint
	month uint
	day   uint
}

type Pesel struct {
	number      string
	gender      string
	dateOfBirth Date
}

type peselError struct {
	number string
}

func (pe peselError) Error() string {
	return fmt.Sprintf("Invalid PESEL: %s", pe.number)
}

func New(number string) (Pesel, error) {
	var pesel Pesel
	peselError := peselError{number: number}

	if len(number) != 11 {
		return pesel, peselError
	}

	return pesel, peselError
}
