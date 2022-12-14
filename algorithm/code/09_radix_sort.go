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

func BitSort(arr []int, bit int) []int {
	length := len(arr)            // 数组长度
	bit_counts := make([]int, 10) // 统计长度 0,1,2,3,4,5,6,7,8,9
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 // 分层处理, bit=1000的三位数不参与排序了，bit=10000的四位数不参与排序了
		bit_counts[num]++          // 统计余数相等的个数
	}
	fmt.Printf("bit_counts: %v\n", bit_counts)

	for i := 1; i < 10; i++ {
		bit_counts[i] += bit_counts[i-1] // 叠加,计算位置
	}
	fmt.Printf("bit_counts: %v\n", bit_counts)

	tmp := make([]int, 10)
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bit_counts[num]-1] = arr[i] // 计算排序的位置
		bit_counts[num]--
	}

	for i := 0; i < length; i++ {
		arr[i] = tmp[i] // 保存数组
	}
	return arr
}

func RadixSort(arr []int) []int {
	max := SelectMax(arr) // 寻找数组极大值
	for bit := 1; max/bit > 0; bit *= 10 {
		// 按照数量级分段
		arr = BitSort(arr, bit) // 每次处理一个级别的排序
		fmt.Printf("arr: %v\n", arr)
	}
	return arr
}

func main() {
	arr := []int{11, 91, 222, 878, 348, 7123, 4213, 6232, 5123, 1011}
	fmt.Printf("RadixSort(arr): %v\n", RadixSort(arr))
}
