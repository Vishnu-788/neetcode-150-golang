package groupanagrams

import (
	"fmt"
	"reflect"
	"testing"
)

func throwErrIfFails(strs []string, expected, actual [][]string, t *testing.T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Error; Expected: %v Actual: %v, Arr: %v\n", expected, actual, strs)
	} else {
		fmt.Print("Success!\n")
	}
}

func TestGrpAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	actual := grpAnagrams(strs)
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	throwErrIfFails(strs, expected, actual, t)

	strs = []string{""}
	actual = grpAnagrams(strs)
	expected = [][]string{{""}}
	throwErrIfFails(strs, expected, actual, t)

	strs = []string{"a"}
	actual = grpAnagrams(strs)
	expected = [][]string{{"a"}}
	throwErrIfFails(strs, expected, actual, t)

}
