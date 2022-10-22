/*
    逐次探索（O(n)）
  		- 前から順に探索するアルゴリズム
*/

package main

import "fmt"

func SequentialSearch(nums []int, target int) bool{
    for i := 0; i < len(nums); i++ {
        if nums[i] == target{
			return true
		}
    }
	return false
}

func Judge(ans bool){
	if ans == true {
		fmt.Println( "Hit" )
	} else {
		fmt.Println( "None" )
	}
}



func main() {
    numbers := []int{7, 65, 8, 32, 4, 21, 10}

    Judge( SequentialSearch(numbers, 5) )
	Judge( SequentialSearch(numbers, 10) )
}