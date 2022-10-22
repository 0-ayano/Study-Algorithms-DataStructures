/*
    ヒープソート（O(n log n)）
  		- クイックソートの次に優秀な整列アルゴリズム
*/

package main

import "fmt"

func HeapSort(numbers []int) []int {
    for i := len(numbers)/2 - 1; i >= 0; i-- {
        heapify(numbers, len(numbers), i)
    }
    for i := len(numbers) - 1; i > 0; i-- {
        numbers[i], numbers[0] = numbers[0], numbers[i]
        heapify(numbers, i, 0)
    }
    return numbers
}

func heapify(numbers []int, len, root int) {
    largest := root
    left := 2*root + 1
    right := 2*root + 2

    if left < len && numbers[left] > numbers[largest] {
        largest = left
    }
    if right < len && numbers[right] > numbers[largest] {
        largest = right
    }
    if largest != root {
        numbers[root], numbers[largest] = numbers[largest], numbers[root] 
        heapify(numbers, len, largest)
    }
}

func main() {
    numbers := []int{7, 65, 8, 32, 4, 21, 10}
    fmt.Println( HeapSort(numbers) )
}