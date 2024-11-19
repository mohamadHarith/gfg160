package main

import (
	"fmt"
)

// Next Permutation

// Given an array of integers arr[] representing a permutation, implement the next permutation that rearranges the numbers into the lexicographically next greater permutation. If no such permutation exists, rearrange the numbers into the lowest possible order (i.e., sorted in ascending order).

// Note - A permutation of n numbers is any possible arrangement of all the integers in range [1-n] where each integer occurs exactly once.

// Examples:

// Input: arr = [2, 4, 1, 7, 5, 0]
// Output: [2, 4, 5, 0, 1, 7]
// Explanation: The next permutation of the given array is {2, 4, 5, 0, 1, 7}.

// Input: arr = [3, 2, 1]
// Output: [1, 2, 3]
// Explanation: As arr[] is the last permutation, the next permutation is the lowest one.

// Input: arr = [3, 4, 2, 5, 1]
// Output: [3, 4, 5, 1, 2]
// Explanation: The next permutation of the given array is {3, 4, 5, 1, 2}.

// Input: arr = [1, 2, 3]
// No of permutations: 3 x 2 x 1 = 6
// Permutation 1: [1,3,2]
// Permutation 2: [2,1,3]
// Permutation 4: [2,3,1]
// Permutation 5: [3,1,2]
// Permutation 6: [3,2,1]

// Input: arr = [1, 2, 3, 4]
// No of permutations: 4 x 3 x 2 x 1 = 24
// Permutation 1: [1,2,4,3]
// Permutation 2: [1,3,2,4]
// Permutation 3: [1,3,4,2]
// Permutation 4: [1,4,2,3]
// Permutation 5: [1,4,3,2]
// so on...

// Constraints:
// 1 ≤ arr.size() ≤ 105
// 1 ≤ arr[i] ≤ 105

// [1,4,3,2,5] [5,2,3,4,1]

func reverse(arr []uint32) {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
}

// NextPermutation:
// 1. find rightmost pivot a[i] < a[i+1]
// 2. find successor a[i] > pivot
// 3. switch successor and pivot
// 4. revers pivot+1 to n
func NextPermutation(arr []uint32) {
	pivotIdx, successorIdx := int32(-1), int32(-1)

	for i := int32(len(arr) - 2); i >= 0; i-- {
		if arr[i] < arr[i+1] {
			pivotIdx = i
			break
		}
	}

	// if no, pivot found, the arr is highest permutation,
	// so reverse it to return to the lowest permutation
	if pivotIdx == -1 {
		reverse(arr)
		return
	}

	for i := int32(len(arr) - 1); i > 0; i-- {
		if arr[i] > arr[pivotIdx] {
			successorIdx = i
			break
		}
	}

	temp := arr[successorIdx]
	arr[successorIdx] = arr[pivotIdx]
	arr[pivotIdx] = temp

	reverse(arr[pivotIdx+1:])

	fmt.Println("after reverse=> ", arr)

}
