/*
    選択ソート（O(n^2)）
  		- ある程度の早さが保証されている整列アルゴリズム
*/

package main

import "fmt"

func SelectionSort(nums []int) []int {
    for i := range nums {
        min_idx := i

        for j := i; j < len(nums); j++ {
            if nums[min_idx] > nums[j] {
                min_idx = j
            }
        }
        nums[i], nums[min_idx] = nums[min_idx], nums[i]
    }
    return nums
}

func main() {
    nums := []int{2, 4, 3, 1, 5, 7, 6}
    fmt.Println(SelectionSort(nums)) // [1 2 3 4 5 6 7]
}