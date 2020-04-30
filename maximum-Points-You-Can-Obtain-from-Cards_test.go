// There are several cards arranged in a row, and each card has an associated number of points The points are given in the integer array cardPoints.
// In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
// Your score is the sum of the points of the cards you have taken.
// Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
// Example 1:
// Input: cardPoints = [1,2,3,4,5,6,1], k = 3
// Output: 12
// Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
// Example 2:
// Input: cardPoints = [2,2,2], k = 2
// Output: 4
// Explanation: Regardless of which two cards you take, your score will always be 4.
// Example 3:
// Input: cardPoints = [9,7,7,9,7,7,9], k = 7
// Output: 55
// Explanation: You have to take all the cards. Your score is the sum of points of all cards.
// Example 4:
// Input: cardPoints = [1,1000,1], k = 1
// Output: 1
// Explanation: You cannot take the card in the middle. Your best score is 1.
// Example 5:
// Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
// Output: 202
// Constraints:
// 1 <= cardPoints.length <= 10^5
// 1 <= cardPoints[i] <= 10^4
// 1 <= k <= cardPoints.length

package main

import "testing"

func Test_maxPoint(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
				k:          3,
			},
			want: 12,
		},
		{
			name: "case 2",
			args: args{
				cardPoints: []int{2, 2, 2},
				k:          2,
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
				k:          3,
			},
			want: 202,
		},
		{
			name: "case 4",
			args: args{
				cardPoints: []int{1, 1000, 1},
				k:          1,
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				cardPoints: []int{9, 7, 7, 9, 7, 7, 9},
				k:          7,
			},
			want: 55,
		},
		{
			name: "case 6",
			args: args{
				cardPoints: []int{99, 82, 25, 56, 39, 77, 22, 58, 64, 77, 19, 36, 93, 14, 19, 12, 94, 76, 93, 24, 92, 67, 18, 37, 37, 60, 87, 28, 64, 7, 29},
				k:          27,
			},
			want: 1477,
		},
		{
			name: "case 7",
			args: args{
				cardPoints: []int{90, 51, 34, 9, 94, 38, 2, 9, 34, 68, 19, 77, 74, 62, 55, 17},
				k:          1,
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoint(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("%v maxPoint() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
