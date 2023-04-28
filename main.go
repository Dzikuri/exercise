package main

import (
	"fmt"

	binarysearch "github.com/dzikuri/exercise/leetcode/binarySearch"
)

func main() {
	// got := romanToInteger.RomanToInteger("IV")
	// fmt.Println(got)

	got := binarysearch.BinarySearch([]int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	fmt.Println(got)
}
