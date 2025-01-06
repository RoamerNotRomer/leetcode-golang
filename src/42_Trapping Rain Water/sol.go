package sol

func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	ans := 0

	leftMax[0] = 0
	rightMax[len(height)-1] = 0

	for i := 1; i <= len(height)-1; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}

	for i := 0; i < len(height); i++ {
		minBorder := min(leftMax[i], rightMax[i])
		if minBorder > height[i] {
			ans = ans + (minBorder - height[i])
		}
	}

	return ans
}
