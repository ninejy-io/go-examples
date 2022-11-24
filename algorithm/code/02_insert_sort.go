package main

import "fmt"

func Insert(arr []int) []int {
	backup := arr[2]
	j := 2 - 1 // 上一个位置循环找到位置插入
	for j >= 0 && backup < arr[j] {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = backup
	fmt.Printf("arr: %v\n", arr)
	return arr
}

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ { // 跳过第一个
			backup := arr[i] // 备份插入的数据
			j := i - 1       // 上一个位置循环找到位置插入
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j] // 从前往后移动
				j--
			}
			arr[j+1] = backup
			fmt.Printf("arr: %v\n", arr)
		}
		return arr
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	// Insert(arr)
	fmt.Printf("InsertSort(arr): %v\n", InsertSort(arr))
}
