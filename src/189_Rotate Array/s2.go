package sol189

func rotate(nums []int, k int) {
	nums1 := make([]int, 0)

	for i := (len(nums) - k%len(nums)); i < len(nums); i++ {
		nums1 = append(nums1, nums[i])
	}
	for i := len(nums) - k%len(nums) - 1; i >= 0; i-- {
		nums[i+k%len(nums)] = nums[i]
	}
	for i := 0; i < k%len(nums); i++ {
		nums[i] = nums1[i]
	}
}
