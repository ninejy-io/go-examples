package utils

import "sort"

func StrInArray(s string, arr []string) bool {
	sort.Strings(arr)
	index := sort.SearchStrings(arr, s)
	if index < len(arr) && arr[index] == s {
		return true
	}
	return false
}
