package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	l := 10
	got := GenerateRandomString(l)
	fmt.Println(got)
	if len(got) != l {
		t.Errorf("GenerateRandomString() = %v, want %v", len(got), l)
	}

	l = 20
	got = GenerateRandomString(l)
	if len(got) != l {
		t.Errorf("GenerateRandomString() = %v, want %v", len(got), l)
	}

}

func TestMask(t *testing.T) {
	s := "1234567890"
	expected1 := "1*********"
	expected2 := "123*******"
	got := Mask(s, 1)
	if len(got) != len(s) {
		t.Errorf("Mask() = %v, want %v", len(got), len(s))
	}

	if got != expected1 {
		t.Errorf("Mask() = %v, want %v", got, expected1)
	}

	s = "1234567890"
	got = Mask(s, 3)
	if len(got) != len(s) {
		t.Errorf("Mask() = %v, want %v", len(got), len(s))
	}

	if got != expected2 {
		t.Errorf("Mask() = %v, want %v", got, expected2)
	}

}
