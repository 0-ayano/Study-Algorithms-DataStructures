/*
    二分探索（O(log n)）
*/

package main

import "fmt"

func BinarySearch(nums []int, target int) bool{
    var low  int = 0
	var high int = len(nums) - 1

	for low <= high {
		var mid int = (low + high)/2
		if( target < nums[mid] ){
			high = mid - 1
		} else if( target > nums[mid] ){
			low = mid + 1
		} else{
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
    numbers := []int{1, 3, 5, 7, 9, 11}

    Judge( BinarySearch(numbers, 5) )
	Judge( BinarySearch(numbers, 10) )
}