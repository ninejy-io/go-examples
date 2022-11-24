package main

import "fmt"

func merge(l_arr []int, r_arr []int) []int {
	l_index := 0
	r_index := 0
	final_arr := []int{} // 最终的数组
	for l_index < len(l_arr) && r_index < len(r_arr) {
		if l_arr[l_index] < r_arr[r_index] {
			final_arr = append(final_arr, l_arr[l_index])
			l_index++
		} else if l_arr[l_index] > r_arr[r_index] {
			final_arr = append(final_arr, r_arr[r_index])
			r_index++
		} else {
			final_arr = append(final_arr, l_arr[l_index], l_arr[l_index])
			l_index++
			r_index++
		}
	}
	for l_index < len(l_arr) { // 把没有结束的归并过来
		final_arr = append(final_arr, l_arr[l_index])
		l_index++
	}
	for r_index < len(r_arr) { // 把没有结束的归并过来
		final_arr = append(final_arr, r_arr[r_index])
		r_index++
	}
	return final_arr
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		l_arr := MergeSort(arr[:mid])
		r_arr := MergeSort(arr[mid:])

		return merge(l_arr, r_arr)
	}
}

func main() {
	arr := []int{31, 23, 1, 9, 2, 8, 26, 13, 7, 6, 4, 5}
	fmt.Printf("MergeSort(arr): %v\n", MergeSort(arr))
}
