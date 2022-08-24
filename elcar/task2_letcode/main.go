package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 4, 4, 5},[]int{3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 4, 5},[]int{1, 2, 3}))
	fmt.Println(findMedianSortedArrays([]int{2, 2},[]int{2, 2}))
	//
	fmt.Println(findMedianSortedArrays([]int{1, 2, 2, 4, 5},[]int{3}))
	fmt.Println(findMedianSortedArrays([]int{1, 3},[]int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr1, arr2 := nums1, nums2
	half := (len(arr1)+len(arr2))/2
	odd := (len(arr1)+len(arr2))%2

	var arr1left int
	var arr1right int
	var arr2left int
	var arr2right int

	if len(arr2) < len(arr1) {
		arr1, arr2 = arr2, arr1
	}

	l, r := 0, len(arr1) -1
	for l<=r {
		i := (l+r) /2
		j := half - 2 - i

		if i>=0 {
			arr1left = arr1[i]
		} else {
			arr1left = math.MinInt64
		}
		if i+1 <len(arr1) {
			arr1right = arr1[i+1]
		} else {
			arr1right = math.MaxInt64
		}
		if j>=0 {
			arr2left = arr2[j]
		} else {
			arr2left = math.MinInt64
		}
		if j+1<len(arr2) {
			arr2right = arr2[j+1]
		} else {
			arr2right = math.MaxInt64
		}

		if arr1left<=arr2right && arr2left <= arr1right {

			if odd == 0 {
				return (float64(max(arr1left, arr2left)) + float64(min(arr1right, arr2right)))/2
			} else {
				return float64(min(arr1right, arr2right))
			}
		} else if arr1left > arr2right{
			r = i-1
		} else {
			l = i + 1
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
