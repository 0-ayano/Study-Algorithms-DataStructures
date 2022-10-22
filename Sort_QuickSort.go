/*
    クイックソート(O(n^2))
		- 最速の整列アルゴリズム
*/

package main

import (
    "fmt"
)

func partition(nums []int, low int, high int) int {
    i := low - 1
    pivot := nums[high]
    for j := low; j < high; j++ {
        if nums[j] <= pivot {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[i+1], nums[high] = nums[high], nums[i+1]
    return i + 1
}

func QuickSort(nums []int, low int, high int) []int {
    if low < high {
        partition_idx := partition(nums, low, high)
        QuickSort(nums, low, partition_idx-1)
        QuickSort(nums, partition_idx+1, high)
    }
    return nums
}

func main() {
    nums := []int{2, 4, 3, 1, 5, 7, 6}
    low := 0
    high := len(nums) - 1
    fmt.Println(QuickSort(nums, low, high)) // [1 2 3 4 5 6 7]
}