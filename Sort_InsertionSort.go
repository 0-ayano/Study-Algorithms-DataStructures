/*
    挿入ソート（O(n^2)）
        - 要素数が少ない、初めからほぼ整列している状態に使う
*/

package main

import (
    "fmt"
)

func InsertionSort(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            j := i
            temp := nums[i]
            for 0 < j && nums[j-1] > temp {
                nums[j] = nums[j-1]
                j--
            }
            nums[j] = temp
        }
    }
    return nums
}

func main() {
    nums := []int{4, 3, 1, 5, 7, 6, 2}
    fmt.Println(InsertionSort(nums)) // [1 2 3 4 5 6 7]
}