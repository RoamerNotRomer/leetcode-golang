package sol

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var depth int
	ans := make([][]int, 0)
	nodeArr := []*TreeNode{root}
	if root == nil {
		return ans
	}

	for depth = 0; len(nodeArr) != 0; depth++ {
		subAns := make([]int, 0)

		tmp := nodeArr
		nodeArr = nil

		for _, node := range tmp {
			subAns = append(subAns, node.Val)

			if node.Left != nil {
				nodeArr = append(nodeArr, node.Left)
			}
			if node.Right != nil {
				nodeArr = append(nodeArr, node.Right)
			}
		}
		if len(subAns) != 0 {
			if depth%2 != 0 {
				reverseArr(subAns)
			}
			ans = append(ans, subAns)
		}
	}
	return ans
}

func reverseArr(arr []int) {
	mid := int((len(arr) - 1) / 2)
	for i := 0; i <= mid; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
