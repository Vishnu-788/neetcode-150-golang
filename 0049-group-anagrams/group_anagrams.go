package groupanagrams

func grpAnagrams(strs []string) [][]string {
	charCount := map[[26]int][]string{}

	for _, str := range strs {
		count := [26]int{}
		for _, c := range str {
			count[c-'a']++
		}

		charCount[count] = append(charCount[count], str)
	}

	res := [][]string{}
	for _, v := range charCount {
		res = append(res, v)
	}

	return res
}
