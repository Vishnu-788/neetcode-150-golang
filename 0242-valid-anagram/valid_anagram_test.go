package validanagram

import (
	"fmt"
	"testing"
)

func throwErrIfFails(s1, s2 string, expected, actual bool, t *testing.T) {
	if actual != expected {
		t.Errorf("Failed; expected: %v, actual %v. s1: %s, s2: %s\n", expected, actual, s1, s2)
	} else {
		fmt.Print("Success\n")
	}
}

func TestValidAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"

	actual := ValidAnagram(s1, s2)
	expected := true
	throwErrIfFails(s1, s2, expected, actual, t)

	s1 = "rat"
	s2 = "car"

	actual = ValidAnagram(s1, s2)
	expected = false
	throwErrIfFails(s1, s2, expected, actual, t)
}
