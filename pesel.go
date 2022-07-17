package pesel

import (
	"fmt"
	"strconv"
	"time"
)

type Gender string
type PeselNumber string

const (
	Female Gender = "female"
	Male   Gender = "male"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type Pesel struct {
	number      PeselNumber
	gender      Gender
	dateOfBirth Date
}

type peselError struct {
	number string
	reason string
}

func (pe peselError) Error() string {
	return fmt.Sprintf("Invalid PESEL : %s. %s", pe.number, pe.reason)
}

func toSlice(number string) ([]int, error) {
	peselNumbers := make([]int, 11)
	var parseError error = nil
	for i := 0; i < 11; i++ {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			parseError = err
			break
		}
		peselNumbers[i] = digit
	}

	return peselNumbers, parseError
}

var weights = [10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}

func calculateChecksum(peselNumbers []int) int {
	var sum int

	for i, digit := range peselNumbers[:10] {
		sum += digit * weights[i]
	}

	digit := sum % 10

	if digit == 0 {
		return 0
	}
	return 10 - digit
}

func decodeGender(peselNumbers []int) Gender {
	if peselNumbers[9]%2 == 0 {
		return Female
	} else {
		return Male
	}
}

var centuries = [5]int{1900, 2000, 2100, 2200, 1800}

func decodeDateOfBirth(peselNumbers []int) (Date, error) {
	var birthDate = &Date{}
	month := 10*peselNumbers[2] + peselNumbers[3]
	mod := month / 20
	date, err := time.Parse("20060102", fmt.Sprintf("%04d%02d%02d", centuries[mod]+10*peselNumbers[0]+peselNumbers[1], month-mod*20, 10*peselNumbers[4]+peselNumbers[5]))

	if err != nil {
		return *birthDate, err
	}

	birthDate.Year = date.Year()
	birthDate.Month = date.Month()
	birthDate.Day = date.Day()

	return *birthDate, nil
}

func New(number string) (Pesel, error) {
	var pesel = &Pesel{}

	if len(number) != 11 {
		return *pesel, peselError{number: number, reason: "Invalid length"}
	}

	peselNumbers, err := toSlice(number)
	if err != nil {
		return *pesel, peselError{number: number, reason: "Characters are invalid"}
	}

	checksum := calculateChecksum(peselNumbers)

	if checksum != peselNumbers[10] {
		return *pesel, peselError{number: number, reason: "Checksum does not match"}
	}

	dateOfBirth, dateErr := decodeDateOfBirth(peselNumbers)
	if dateErr != nil {
		return *pesel, peselError{number: number, reason: "Encoded date is invalid"}
	}

	pesel.number = PeselNumber(number)
	pesel.gender = decodeGender(peselNumbers)
	pesel.dateOfBirth = dateOfBirth

	return *pesel, nil
}
