package pesel

type Date struct {
	year  uint
	month uint
	day   uint
}

type Pesel struct {
	value       string
	gender      string
	dateOfBirth Date
}
