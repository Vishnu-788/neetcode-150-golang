package validanagram

func ValidAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freqArr := [26]int{}

	for i := 0; i < len(s1); i++ {
		freqArr[s1[i]-'a']++
		freqArr[s2[i]-'a']--
	}

	for _, v := range freqArr {
		if v != 0 {
			return false
		}
	}
	return true
}
