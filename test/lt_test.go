package test

import (
	"sort"
	"testing"
)

//字母异位
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
func Test_main(t *testing.T) {
	lengthOfLongestSubstring("abcabcbb")
}

func groupAnagrams(strs []string) [][]string {
	//声明一个hash存储排序后的字符串
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortStr := string(s)
		mp[sortStr] = append(mp[sortStr], str)
	}
	res := [][]string{}
	for _, str := range mp {
		res = append(res, str)
	}
	return res
}

func longestConsecutive(nums []int) int {
	mp := map[int]bool{}
	for _, num := range nums {
		mp[num] = true
	}
	longestStreak := 0
	for num := range mp {
		if !mp[num-1] {
			currentNum := num
			currentStreak := 1
			for mp[currentNum+1] {
				currentNum++
				currentStreak++
			}
			longestStreak = max(longestStreak, currentStreak)
		}
	}
	return longestStreak
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//移动零
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

//盛最多水的容器
func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height)
	for left < right {
		heightLeft := height[left]
		heightRight := height[right]
		tempArea := (right - left) * min(heightRight, heightLeft)
		maxArea = max(maxArea, tempArea)
	}
	return maxArea
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

//三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := length - 1
		for L < R {
			temp := nums[i] + nums[L] + nums[R]
			if temp == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if temp > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}

//无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	//if len(s) == 0 {
	//	return 0
	//}
	//mp := map[byte]int{}
	//maxLength := 0
	//left := 0
	//for right := 0; right < len(s); right++ {
	//	if index, ok := mp[s[right]]; ok {
	//		left = max(left, index+1)
	//		fmt.Println(index, left)
	//	}
	//	mp[s[right]] = right
	//	maxLength = max(maxLength, right-left+1)
	//}
	//return maxLength
	/*
		滑动窗口
	*/

	window := map[byte]int{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

//找到字符串中所有字母异位词
func findAnagrams(s, t string) []int {
	need := map[byte]int{}
	window := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right, valid := 0, 0, 0
	var res []int
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//判断左侧窗口是否要收缩
		for right-left >= len(t) {
			//当窗口符合条件时，把起始索引加入结果
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//相交链表
func getIntersectionNode(heaA, headB *ListNode) *ListNode {
	if heaA == nil || headB == nil {
		return nil
	}
	pa, pb := heaA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = heaA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
