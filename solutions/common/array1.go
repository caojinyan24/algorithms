package common

import (
	"fmt"
	"sort"
)

func RunMinimalHeaviestSetA() []int32 {
	// return minimalHeaviestSetA([]int32{5, 3, 2, 4, 1, 2})
	return minimalHeaviestSetA([]int32{1, 1, 1, 1, 1, 1, 1, 1})

}
func minimalHeaviestSetA(arr []int32) []int32 {
	// Write your code here
	//so len(A)<len(B)&&sum(A)>sum(B)
	//sum all arr
	//sort
	//
	if len(arr) <= 1 {
		return arr
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println("arr=", arr)
	var sum int32 = 0
	for _, item := range arr {
		sum = sum + item
	}
	var result = make([]int32, 0)
	var resultSum int32 = 0
	for i := len(arr) - 1; i >= 0; i-- {
		resultSum = resultSum + arr[i]
		result = append([]int32{arr[i]}, result...)
		if resultSum > sum-resultSum {
			return result
		}
	}
	return []int32{}
}
