package main

// You are given an array of integer arr[] where each number represents a vote to a candidate. Return the candidates that have votes greater than one-third of the total votes, If there's not a majority vote, return an empty array.

// Note: The answer should be returned in an increasing format.

// Examples:

// Input: arr[] = [2, 1, 5, 5, 5, 5, 6, 6, 6, 6, 6]
// Output: [5, 6]
// Explanation: 5 and 6 occur more n/3 times.

// Input: arr[] = [1, 2, 3, 4, 5]
// Output: []
// Explanation: no candidate occur more than n/3 times.
// Constraint:
// 1 <= arr.size() <= 106
// -109 <= arr[i] <= 109

func MajorityElement(nums []int) []int {
	var count1, count2 int
	var element1 *int
	var element2 *int
	var res []int

	for i := 0; i < len(nums); i++ {
		v := nums[i]

		if count1 == 0 || element1 == nil {
			if element2 != nil && *element2 == v {
				// skip
			} else {
				element1 = &v
			}
		}

		if count2 == 0 || element2 == nil {
			if element1 != nil && *element1 == v {
				// skip
			} else {
				element2 = &v
			}
		}

		if v == *element1 {
			count1++
		} else if v == *element2 {
			count2++
		} else {
			if count1 > 0 {
				count1--
			}
			if count2 > 0 {
				count2--
			}
		}

	}

	count1 = 0
	count2 = 0
	for _, v := range nums {
		if element1 != nil && *element1 == v {
			count1++
		} else if element2 != nil && *element2 == v {
			count2++
		}
	}

	if count1 > len(nums)/3 && count2 > len(nums)/3 {
		if *element1 > *element2 {
			res = append(res, *element2, *element1)
		} else {
			res = append(res, *element1, *element2)
		}
	} else if count1 > len(nums)/3 {
		res = append(res, *element1)
	} else if count2 > len(nums)/3 {
		res = append(res, *element2)
	}

	return res
}
