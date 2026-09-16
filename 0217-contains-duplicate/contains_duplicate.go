package contains_duplicate

// nums = [1, 2, 3, 1]

func ContainsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, exists := set[v]; exists {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
