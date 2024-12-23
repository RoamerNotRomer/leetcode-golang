package sol350

func partition(nums []int, low, high int) int {
	pivot := nums[high]

	smallerIdx := low

	for i := low; i <= high-1; i++ {
		if nums[i] < pivot {
			nums[i], nums[smallerIdx] = nums[smallerIdx], nums[i]
			smallerIdx++
		}
	}

	nums[smallerIdx], nums[high] = nums[high], nums[smallerIdx]
	return smallerIdx
}

func qsort_iteration(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)

	qsort_iteration(nums, low, p-1)
	qsort_iteration(nums, p+1, high)
}

func intersect(nums1 []int, nums2 []int) []int {
	qsort_iteration(nums1, 0, len(nums1)-1)
	qsort_iteration(nums2, 0, len(nums2)-1)

	intersect := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			intersect = append(intersect, nums1[i])
			i++
			j++
		}
	}

	return intersect
}

func main() {
	test := []int{5, 9, 3, 1, 2, 4, 7, 10, 6}

	qsort_iteration(test, 0, 8)

	for _, v := range test {
		println("%d\n", v)
	}

}
