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

import "fmt"

func maxPoint(cardPoints []int, k int) int {
	i, j, scoreI, scoreJ, score := 0, len(cardPoints)-1, 0, 0, 0

	fmt.Println(cardPoints, k)
	j = len(cardPoints) - i
	for i = 0; i < k/2; i++ {
		j = len(cardPoints) - i - 1
		scoreI += cardPoints[i]
		scoreJ += cardPoints[j]
		fmt.Println(i, j, scoreI, scoreJ, scoreI+scoreJ)
	}

	if len(cardPoints) == k {
		return scoreI + scoreJ + cardPoints[i]
	}
	adjI := 0
	if k%2 != 0 {
		if cardPoints[i] > cardPoints[j-1] {
			scoreI += cardPoints[i]
			adjI++
			i++
		} else {
			scoreJ += cardPoints[j-1]
			j--
		}
	}
	score = scoreI + scoreJ
	scoreI, scoreJ = score, score

	fmt.Println(i-1, j, scoreI, scoreJ, score)
	ii := i - 1
	jj := j
	for ; i < k; i++ {
		j--
		if ii < 0 || jj >= len(cardPoints) {
			break
		}

		scoreI += cardPoints[i] - cardPoints[jj]
		scoreJ += cardPoints[j] - cardPoints[ii]

		score = max(score, max(scoreI, scoreJ))
		fmt.Println(i, j, ii, jj, scoreI, scoreJ, score)
		ii--
		jj++
	}
	if k > 1 && k%2 != 0 {
		adjII, adjJJ := 0, 0
		if ii >= 0 {
			adjII = cardPoints[ii]
		}
		if jj < len(cardPoints) {
			adjJJ = cardPoints[jj]
		}
		if adjI > 0 {
			scoreJ += cardPoints[j] - adjII
		} else {
			scoreI += cardPoints[i] - adjJJ
		}
		score = max(score, max(scoreI, scoreJ))
	}
	fmt.Println(i, j, ii, jj, scoreI, scoreJ, score)

	return score
}
