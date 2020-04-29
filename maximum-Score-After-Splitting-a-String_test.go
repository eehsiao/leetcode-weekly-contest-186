// Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).
// The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.
// Example 1:
// Input: s = "011101"
// Output: 5
// Explanation:
// All possible ways of splitting s into two non-empty substrings are:
// left = "0" and right = "11101", score = 1 + 4 = 5
// left = "01" and right = "1101", score = 1 + 3 = 4
// left = "011" and right = "101", score = 1 + 2 = 3
// left = "0111" and right = "01", score = 1 + 1 = 2
// left = "01110" and right = "1", score = 2 + 1 = 3
// Example 2:
// Input: s = "00111"
// Output: 5
// Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
// Example 3:
// Input: s = "1111"
// Output: 3
// Constraints:
// 2 <= s.length <= 500
// The string s consists of characters '0' and '1' only.

package main

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "011101",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "1111",
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				s: "0000",
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				s: "00111",
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				s: "11100",
			},
			want: 2,
		},
		{
			name: "case 6",
			args: args{
				s: "0100110001011111011011010010001111000100110011101110001001111010010011001111000",
			},
			want: 44,
		},
		{
			name: "case 7",
			args: args{
				s: "01001",
			},
			want: 4,
		},
		{
			name: "case 8",
			args: args{
				s: "010",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("%v maxScore() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
