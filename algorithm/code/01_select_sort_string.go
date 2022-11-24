package main

import (
	"fmt"
	"strings"
)

func StringCompare() {
	// fmt.Println("a" < "b") // 字符串存在地址的比较
	pa := "a"
	pb := "b"
	pa2 := "a"
	fmt.Printf("pa: %v\n", &pa)
	fmt.Printf("pb: %v\n", &pb)
	fmt.Printf("pa2: %v\n", &pa2)
	// a<b<c 首先比较第一个字母, 左边小于右边 -1, 左边大于右边 1, 相等 0
	// 第一个字母比较不成功比较第二个
	fmt.Printf("strings.Compare(\"a\", \"b\"): %v\n", strings.Compare("a", "b"))
	fmt.Printf("strings.Compare(\"b\", \"a\"): %v\n", strings.Compare("b", "a"))
	fmt.Printf("strings.Compare(\"a\", \"a\"): %v\n", strings.Compare("a", "a"))
	fmt.Printf("strings.Compare(\"ab\", \"ac\"): %v\n", strings.Compare("ab", "ac"))
}

func SelectSortMaxString(arr []string) string {
	length := len(arr)
	if length <= 0 {
		return arr[0]
	} else {
		max := arr[0] // 假定第一个最大
		for i := 1; i < length; i++ {
			if strings.Compare(max, arr[i]) < 0 { // 任何一个比我大的数，最大的
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSortString(arr []string) []string {
	length := len(arr)
	if length <= 0 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { // 只剩一个元素不需要挑选
			min := i                          // 标记索引
			for j := i + 1; j < length; j++ { // 每次选出一个极小值
				if strings.Compare(arr[min], arr[j]) > 0 {
					min = j // 保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] // 数据交换
			}
			fmt.Printf("arr: %v\n", arr)
		}
		return arr
	}
}

func main() {
	// StringCompare()
	arr := []string{"d", "c", "x", "p", "m", "a", "y", "j", "z"}
	// fmt.Printf("SelectSortMaxString(arr): %v\n", SelectSortMaxString(arr))
	fmt.Printf("SelectSortString(arr): %v\n", SelectSortString(arr))
}
