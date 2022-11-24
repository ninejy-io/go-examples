package main

import "fmt"

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { // 插入排序的变种
		backup := arr[i] // 备份插入的数据
		j := i - gap     // 上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}
		arr[j+gap] = backup // 插入
		fmt.Printf("arr: %v\n", arr)
	}
}

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { // 处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			// gap--
			gap /= 2
		}
		return arr
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	fmt.Printf("ShellSort(arr): %v\n", ShellSort(arr))
}
