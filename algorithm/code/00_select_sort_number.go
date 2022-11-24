package main

import "fmt"

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0] // 假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max { // 任何一个比我大的数，最大
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { // 只剩一个元素不需要挑选
			min := i //标记索引
			for j := i + 1; j < length; j++ {
				if arr[min] < arr[j] {
					min = j // 保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] // 数据交换
			}
			// fmt.Printf("arr: %v\n", arr)
		}
		return arr
	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 0}
	// fmt.Printf("SelectSortMax(arr): %v\n", SelectSortMax(arr))
	fmt.Printf("SelectSort(arr): %v\n", SelectSort(arr))
}
