package leetcode28

func strStr(haystack string, needle string) int {
	nLen := len(needle)
	lps := make([]int, nLen)
	for i := range needle {
		if i == 0 {
			lps[0] = 0
			continue
		}

		lpsN := lps[i-1]

		for lpsN > 0 && needle[i] != needle[lpsN] {
			lpsN = lps[lpsN-1]
		}

		if needle[i] == needle[lpsN] {
			lpsN++
		}
		lps[i] = lpsN
	}

	i := 0
	j := 0

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}

		if j == len(needle) {
			return i - j
		}
	}
	return -1

}
