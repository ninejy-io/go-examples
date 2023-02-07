package main

import "fmt"

func CocktailSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length/2; i++ { // 每次循环，正向冒泡一次，反向冒泡一次
		left := 0           // 左边
		right := length - 1 // 右边

		for left <= right { // 循环结束的条件
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++

			if arr[right-1] > arr[right] {
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
		fmt.Printf("arr: %v\n", arr)
	}
	return arr
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	fmt.Printf("CocktailSort(arr): %v\n", CocktailSort(arr))
}
