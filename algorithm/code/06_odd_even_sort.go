package main

import "fmt"

func OddEvenSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		is_sorted := false
		for is_sorted == false {
			is_sorted = true
			for i := 1; i < length-1; i = i + 2 { // 奇数位
				if arr[i] > arr[i+1] { // 需要交换就换
					arr[i], arr[i+1] = arr[i+1], arr[i]
					is_sorted = false
				}
			}
			fmt.Printf("arr---1: %v\n", arr)
			for i := 0; i < length-1; i = i + 2 { // 偶数位
				if arr[i] > arr[i+1] { // 需要交换就换
					arr[i], arr[i+1] = arr[i+1], arr[i]
					is_sorted = false
				}
			}
			fmt.Printf("arr---0: %v\n", arr)
		}
		return arr
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	fmt.Printf("OddEvenSort(arr): %v\n", OddEvenSort(arr))
}
