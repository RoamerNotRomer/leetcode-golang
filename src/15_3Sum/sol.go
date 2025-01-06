package sol

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		//remove duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		target := -1 * nums[i]

		for j := i + 1; j < n; j++ {
			//remove duplicate
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
