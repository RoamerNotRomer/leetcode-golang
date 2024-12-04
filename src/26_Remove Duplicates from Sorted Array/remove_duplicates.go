package removeduplicates

func removeDuplicates(nums []int) int {

	piv := nums[0]
	pivIdx := 0
	nums[pivIdx] = piv
	pivIdx++

	for i := 1; i < len(nums); i++ {
		if piv != nums[i] {
			piv = nums[i]
			nums[pivIdx] = piv
			pivIdx++
		}
	}

	nums = nums[0:pivIdx]

	return len(nums)
}
