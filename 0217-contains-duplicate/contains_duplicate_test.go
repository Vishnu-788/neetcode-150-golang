package contains_duplicate

import (
	"fmt"
	"testing"
)

func throwErrIfFails(nums []int, expected, actual bool, t *testing.T) {
	if actual != expected {
		t.Errorf("Test case failed: Expected: %v Actual: %v. Input: %v\n", expected, actual, nums)
	} else {
		fmt.Printf("Test Passed\n")
	}
}

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	actual := ContainsDuplicate(nums)
	expected := true

	throwErrIfFails(nums, expected, actual, t)

	nums = []int{1, 2, 3, 4}
	actual = ContainsDuplicate(nums)
	expected = false

	throwErrIfFails(nums, expected, actual, t)

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	actual = ContainsDuplicate(nums)
	expected = true

	throwErrIfFails(nums, expected, actual, t)
}
