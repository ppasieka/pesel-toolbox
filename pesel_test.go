package pesel

import (
	"testing"
)

func TestPeselInvalidLength(t *testing.T) {
	cases := []string{
		"0628255417",
		"06",
		"",
		"0628255417312",
	}

	for _, data := range cases {
		_, err := New(data)
		if err == nil {
			t.Errorf("PESEL %s length is invalid", data)
		}
	}
}

type checksumTestData struct {
	pesel            string
	expectedChecksum int
}

func TestChecksumCalculation(t *testing.T) {
	cases := []checksumTestData{
		{pesel: "06282554176", expectedChecksum: 6},
		{pesel: "95012037841", expectedChecksum: 1},
		{pesel: "50052453448", expectedChecksum: 8},
		{pesel: "57081443219", expectedChecksum: 9},
		{pesel: "89092243964", expectedChecksum: 4},
		{pesel: "59090217216", expectedChecksum: 6},
		{pesel: "02242225878", expectedChecksum: 8},
		{pesel: "49111167646", expectedChecksum: 6},
		{pesel: "03211217445", expectedChecksum: 5},
		{pesel: "85091728442", expectedChecksum: 2},
	}

	for _, data := range cases {
		number, _ := toSlice(data.pesel)
		checksum := calculateChecksum(number)
		if checksum != data.expectedChecksum {
			t.Errorf("Invalid checksum for PESEL %s. Expected %d, but got %d", data.pesel, data.expectedChecksum, checksum)
		}
	}
}

func TestValidPeselNumbers(t *testing.T) {
	cases := []string{
		"06282554176",
		"95012037841",
		"50052453448",
		"57081443219",
		"89092243964",
		"59090217216",
		"02242225878",
		"49111167646",
		"03211217445",
		"85091728442",
	}
	for _, data := range cases {
		_, err := New(data)
		if err != nil {
			t.Errorf("PESEL %s should be valid", data)
		}
	}
}

func TestInvalidPeselNumbers(t *testing.T) {
	cases := []string{
		"12345678403",
	}
	for _, data := range cases {
		_, err := New(data)
		if err == nil {
			t.Errorf("PESEL %s should be invalid", data)
		}
	}
}

type genderTestData struct {
	pesel          string
	expectedGender Gender
}

func TestGenderIsEncoded(t *testing.T) {
	cases := []genderTestData{
		{pesel: "06282554176", expectedGender: "male"},
		{pesel: "95012037841", expectedGender: "female"},
		{pesel: "50052453448", expectedGender: "female"},
		{pesel: "57081443219", expectedGender: "male"},
		{pesel: "89092243964", expectedGender: "female"},
		{pesel: "59090217216", expectedGender: "male"},
		{pesel: "02242225878", expectedGender: "male"},
		{pesel: "49111167646", expectedGender: "female"},
		{pesel: "03211217445", expectedGender: "female"},
		{pesel: "85091728442", expectedGender: "female"},
	}

	for _, data := range cases {
		pesel, _ := New(data.pesel)
		if pesel.gender != data.expectedGender {
			t.Errorf("Invalid  for PESEL %s. Expected %s, but got %s", data.pesel, data.expectedGender, pesel.gender)
		}
	}
}
