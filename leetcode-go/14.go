package main

func longestCommonPrefix(strs []string) string {
	// Use shortest string as baseline match
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				// 0 - i in array
				return strs[j][:i]
			}
		}
	}

	return short
}

// Find shortest string
func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// set initial to 0
	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}
