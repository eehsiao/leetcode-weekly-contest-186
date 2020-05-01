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

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 4, 2, 7, 5, 3, 8, 6, 9},
		},
		{
			name: "case 2",
			args: args{
				nums: [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}},
			},
			want: []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16},
		},
		{
			name: "case 3",
			args: args{
				nums: [][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}},
			},
			want: []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
