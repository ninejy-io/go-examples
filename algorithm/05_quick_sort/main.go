package main

import "fmt"

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		split_data := arr[0]          // 以第一个为基准
		low := make([]int, 0, 0)      // 存储比我小的
		high := make([]int, 0, 0)     // 存储比我大的
		mid := make([]int, 0, 0)      // 存储与我相等的
		mid = append(mid, split_data) // 加入第一个相等的
		for i := 1; i < length; i++ {
			if arr[i] < split_data {
				low = append(low, arr[i])
			} else if arr[i] > split_data {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) // 切割递归处理
		_arr := append(append(low, mid...), high...)
		return _arr
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	fmt.Printf("QuickSort(arr): %v\n", QuickSort(arr))
}
