package sol

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target <= nums[mid] && nums[0] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target >= nums[mid] && nums[len(nums)-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
