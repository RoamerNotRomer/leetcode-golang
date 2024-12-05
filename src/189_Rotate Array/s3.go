package sol189

func rotate(nums []int, k int) {
	reverse(nums[:len(nums)-k%len(nums)])
	reverse(nums[len(nums)-k%len(nums):])
	reverse(nums)
}

func reverse(arr []int) {
	for i := 0; i <= (len(arr)/2 - 1); i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
