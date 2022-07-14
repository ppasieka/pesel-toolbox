package pesel

import (
	"testing"
)

func Test(t *testing.T) {
	_, e := New("123")
	if e == nil {
		t.Error("Should fail when crating a PESEL")
	}
}
