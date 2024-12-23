package main

func plusOne(digits []int) []int {
	res := []int{}
	length := len(digits)
	plusOneMeth(&digits, length)
	newLength := len(digits)

	if newLength > length {
		res = append(res, 1)
	}

	res = append(res, digits[:length]...)
	return res
}

func plusOneMeth(digits *[]int, length int) {
	if length == 0 {
		*digits = append(*digits, 1)
		return
	}
	if (*digits)[length-1] < 9 {
		(*digits)[length-1]++
	} else {
		(*digits)[length-1] = 0
		plusOneMeth(digits, length-1)
	}
	return
}

func main() {
	test := []int{9}
	res := plusOne(test)

	for _, val := range res {
		println(val)
	}
}
