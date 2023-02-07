package main

import "fmt"

func SelectMax(arr []int) int {
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

func CountSort(arr []int) []int {
	max := SelectMax(arr)               // 寻找最大值
	sorted_arr := make([]int, len(arr)) // 排序之后存储
	counts_arr := make([]int, len(arr)) // 统计次数

	for _, v := range arr {
		counts_arr[v]++
	}
	fmt.Printf("第一次统计次数 counts_arr: %v\n", counts_arr)

	for i := 1; i <= max; i++ {
		counts_arr[i] += counts_arr[i-1] // 叠加
	}
	fmt.Printf("次数叠加 counts_arr: %v\n", counts_arr)

	for _, v := range arr {
		sorted_arr[counts_arr[v]-1] = v // 展开数据
		counts_arr[v]--
		fmt.Printf("sorted_arr: %v\n", sorted_arr)
		fmt.Printf("counts_arr: %v\n", counts_arr)
	}

	return sorted_arr
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 5, 5, 3, 4, 3, 2, 1}
	fmt.Printf("CountSort(arr): %v\n", CountSort(arr))
}
