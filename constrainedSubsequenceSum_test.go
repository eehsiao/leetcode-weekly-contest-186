// Given an integer array nums and an integer k, return the maximum sum of a non-empty subsequence of that array such that for every two consecutive integers in the subsequence, nums[i] and nums[j], where i < j, the condition j - i <= k is satisfied.
// A subsequence of an array is obtained by deleting some number of elements (can be zero) from the array, leaving the remaining elements in their original order.
// Example 1:
// Input: nums = [10,2,-10,5,20], k = 2
// Output: 37
// Explanation: The subsequence is [10, 2, 5, 20].
// Example 2:
// Input: nums = [-1,-2,-3], k = 1
// Output: -1
// Explanation: The subsequence must be non-empty, so we choose the largest number.
// Example 3:
// Input: nums = [10,-2,-10,-5,20], k = 2
// Output: 23
// Explanation: The subsequence is [10, -2, -5, 20].
// Constraints:
// 1 <= k <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4

package main

import "testing"

func Test_constrainedSubsetSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{10, 2, -10, 5, 20},
				k:    2,
			},
			want: 37,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{-1, -2, -3},
				k:    1,
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{10, -2, -10, -5, 20},
				k:    2,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constrainedSubsetSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("%v constrainedSubsetSum() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
