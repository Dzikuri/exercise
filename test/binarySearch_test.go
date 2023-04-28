package test

import (
	"testing"

	binarysearch "github.com/dzikuri/exercise/leetcode/binarySearch"
)

type binarySearchTest struct {
	arg1   []int
	target int
	expect int
}

var binarySearchTests = []binarySearchTest{
	binarySearchTest{[]int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 10},
	binarySearchTest{[]int{-2, 1, 23, 3, 4, 5, -1, 10, 12, 13, 11}, 12, 8},
	binarySearchTest{[]int{-3, 1, 25, 3, 4, 5, -1, 10, 12, 23, 11}, 2, -1},
	binarySearchTest{[]int{5, 1, 26, 3, 4, 5, -1, 10, 12, 13, 16}, 1, 1},
	binarySearchTest{[]int{9, 1, 9, 3, 4, 5, -1, 10, 12, 13, 20}, 0, -1},
}

func TestBinarySearch_1(t *testing.T) {
	for _, test := range binarySearchTests {
		if output := binarysearch.BinarySearch(test.arg1, test.target); output != test.expect {
			t.Errorf("Testing Binary Search array : %v target %d expected %d but result %d", test.arg1, test.target, test.expect, output)
		}
	}
}
