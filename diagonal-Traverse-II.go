// Given a list of lists of integers, nums, return all elements of nums in diagonal order as shown in the below images.
// Example 1:
// https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png
// Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,4,2,7,5,3,8,6,9]
// Example 2:
// https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png
// Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
// Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
// Example 3:
// Input: nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
// Output: [1,4,2,5,3,8,6,9,7,10,11]
// Example 4:
// Input: nums = [[1,2,3,4,5,6]]
// Output: [1,2,3,4,5,6]
// Constraints:
// 1 <= nums.length <= 10^5
// 1 <= nums[i].length <= 10^5
// 1 <= nums[i][j] <= 10^9
// There at most 10^5 elements in nums.

package main

func findDiagonalOrder(nums [][]int) []int {
	var (
		n, m int
		ret  []int
	)
	n = len(nums)
	for i := 0; i < n; i++ {
		m = max(m, len(nums[i]))
	}

	for d := 0; d < n; d++ {
		x, y := 0, d

		for x < m && y >= 0 {
			if x < len(nums[y]) {
				ret = append(ret, nums[y][x])
			}
			x++
			y--
		}
	}

	for w := 1; w < m; w++ {
		x, y := w, n-1

		for x < m && y >= 0 {
			if x < len(nums[y]) {
				ret = append(ret, nums[y][x])
			}
			x++
			y--
		}
	}

	return ret
}
