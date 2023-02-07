package main

import "fmt"

func HeapSortMax(arr []int, length int) []int {
	if length <= 0 {
		return arr
	} else {
		depth := length/2 - 1
		for i := depth; i >= 0; i-- { // 循环所有的三节点
			top_max := i // 假定最大的在i的位置
			left_child := 2*i + 1
			right_child := 2*i + 2                                        // 左右孩子节点
			if left_child <= length-1 && arr[left_child] > arr[top_max] { // 防止越界
				top_max = left_child // 如果左边比我大, 记录最大值
			}
			if right_child <= length-1 && arr[right_child] > arr[top_max] {
				top_max = right_child // 如果右边比我大, 记录最大值
			}
			if top_max != i { // 确保i的值就是最大
				arr[i], arr[top_max] = arr[top_max], arr[i]
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		last_mess_len := length - i // 每次截取一段
		HeapSortMax(arr, last_mess_len)
		arr[0], arr[last_mess_len-1] = arr[last_mess_len-1], arr[0]
	}
	return arr
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	// fmt.Printf("HeapSortMax(arr): %v\n", HeapSortMax(arr, len(arr)))
	fmt.Printf("HeapSort(arr): %v\n", HeapSort(arr))
}
