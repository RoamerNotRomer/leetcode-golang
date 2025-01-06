package sol

func lengthOfLongestSubstring1(s string) int {
	var left, right, longest int
	var dir map[string]int
	longest = 0
	left = 0

	for left < len(s) && right < len(s) {
		right = left
		dir = make(map[string]int)
		exist := false
		repeat := ""
		for right < len(s) && (!exist) {
			_, exist = dir[string(s[right])]
			if !exist {
				dir[string(s[right])] = 1
				right++
			} else {
				repeat = string(s[right])
				break
			}
		}
		if longest < (right - left) {
			longest = right - left
		}
		for i := left; i < right; i++ {
			if string(s[i]) == repeat {
				left = i + 1
				break
			}
		}
	}
	return longest
}
