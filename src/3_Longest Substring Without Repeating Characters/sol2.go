package sol

func lengthOfLongestSubstring(s string) int {
	var left, right, longest int
	dir := make(map[string]int)
	longest = 0
	right, left = 0, 0

	for right = 0; right < len(s); right++ {

		for dir[string(s[right])] != 0 {
			dir[string(s[left])]--
			left++
		}

		dir[string(s[right])]++
		longest = max(longest, right-left+1)
	}
	return longest
}

func main() {

	test := "pwwkew"

	println(lengthOfLongestSubstring(test))

}
