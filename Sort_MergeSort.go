/*
    マージソート（O(n log n）
  		- 
*/

package main

import "fmt"

func MergeSort(numbers []int) []int {
    if len(numbers) < 2 {
        return numbers
    }
    middle := len(numbers) / 2
    return merge(MergeSort(numbers[:middle]), MergeSort(numbers[middle:]))
}

func merge(left, right []int) []int {
    var merged []int

    for i := 0; len(left) > 0 && len(right) > 0; i++ {
        if left[0] < right[0] {
            merged = append(merged, left[0])
            left = left[1:]
        } else {
            merged = append(merged, right[0])
            right = right[1:]
        }
    }

    for _, l := range left {
        merged = append(merged, l)
    }
    for _, r := range right {
        merged = append(merged, r)
    }
    return merged
}

func main() {
    numbers := []int{7, 65, 8, 32, 4, 21, 10}
    fmt.Println( MergeSort(numbers) )
}