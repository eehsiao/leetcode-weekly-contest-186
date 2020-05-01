# leetcode-weekly-contest-186
186th LeetCode Weekly Contest

[Challenge site](https://leetcode.com/contest/weekly-contest-186/)

[My leetcode](https://leetcode.com/eehsiao/)

# Problem list

## Maximum Score After Splitting a String
###  [Maximum Score After Splitting a String](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/maximum-Score-After-Splitting-a-String.go) [(Test Case)](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/maximum-Score-After-Splitting-a-String_test.go)
```
Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).

The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.

 

Example 1:

Input: s = "011101"
Output: 5 
Explanation: 
All possible ways of splitting s into two non-empty substrings are:
left = "0" and right = "11101", score = 1 + 4 = 5 
left = "01" and right = "1101", score = 1 + 3 = 4 
left = "011" and right = "101", score = 1 + 2 = 3 
left = "0111" and right = "01", score = 1 + 1 = 2 
left = "01110" and right = "1", score = 2 + 1 = 3
Example 2:

Input: s = "00111"
Output: 5
Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
Example 3:

Input: s = "1111"
Output: 3
 

Constraints:

2 <= s.length <= 500
The string s consists of characters '0' and '1' only.
```

## Maximum Points You Can Obtain from Cards
###  [Maximum Points You Can Obtain from Cards](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/maximum-Points-You-Can-Obtain-from-Cards.go) [(Test Case)](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/maximum-Points-You-Can-Obtain-from-Cards_test.go)
```
There are several cards arranged in a row, and each card has an associated number of points The points are given in the integer array cardPoints.
In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
Your score is the sum of the points of the cards you have taken.
Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
Example 1:
Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
Example 2:
Input: cardPoints = [2,2,2], k = 2
Output: 4
Explanation: Regardless of which two cards you take, your score will always be 4.
Example 3:
Input: cardPoints = [9,7,7,9,7,7,9], k = 7
Output: 55
Explanation: You have to take all the cards. Your score is the sum of points of all cards.
Example 4:
Input: cardPoints = [1,1000,1], k = 1
Output: 1
Explanation: You cannot take the card in the middle. Your best score is 1.
Example 5:
Input: cardPoints = [1,79,80,1,1,1,200,1], k = 3
Output: 202
Constraints:
1 <= cardPoints.length <= 10^5
1 <= cardPoints[i] <= 10^4
1 <= k <= cardPoints.length
```

## Diagonal Traverse II
###  [Diagonal Traverse II](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/diagonal-Traverse-II.go) [(Test Case)](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/diagonal-Traverse-II_test.go)
```
Given a list of lists of integers, nums, return all elements of nums in diagonal order as shown in the below images.
Example 1:
https://assets.leetcode.com/uploads/2020/04/08/sample_1_1784.png
Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,4,2,7,5,3,8,6,9]
Example 2:
https://assets.leetcode.com/uploads/2020/04/08/sample_2_1784.png
Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
Example 3:
Input: nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
Output: [1,4,2,5,3,8,6,9,7,10,11]
Example 4:
Input: nums = [[1,2,3,4,5,6]]
Output: [1,2,3,4,5,6]
Constraints:
1 <= nums.length <= 10^5
1 <= nums[i].length <= 10^5
1 <= nums[i][j] <= 10^9
There at most 10^5 elements in nums.
```

## Constrained Subset Sum
###  [Constrained Subset Sum](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/constrainedSubsequenceSum.go) [(Test Case)](https://github.com/eehsiao/leetcode-weekly-contest-186/blob/master/constrainedSubsequenceSum_test.go)
```
```
