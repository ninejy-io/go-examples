package main

import "fmt"

func BubbleFindMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] { // 两两比较
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		return arr[length-1]
	}
}

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { // 只剩一个，不需要冒泡了
			is_need_exchange := false
			for j := 0; j < length-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					is_need_exchange = true
				}
			}
			if !is_need_exchange {
				break
			}
			fmt.Printf("arr: %v\n", arr)
		}
		return arr
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	// fmt.Printf("BubbleFindMax(arr): %v\n", BubbleFindMax(arr))
	fmt.Printf("BubbleSort(arr): %v\n", BubbleSort(arr))
}
