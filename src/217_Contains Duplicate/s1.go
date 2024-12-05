package sol217

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for _, val := range nums {
		_, ok := dict[val]
		if !ok {
			dict[val]++
		} else {
			return true
		}
	}
	return false
}
