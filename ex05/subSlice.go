package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if length > capacity {
		capacity = length
	}
	ret := make([]int, length, capacity)
	j := 0
	for i := begin; i < len(slice); i++ {
		if j < length {
			ret[j] = slice[i]
			j++
		} else {
			return ret
		}
	}
	return ret
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int

	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
