package main

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {

	maxLen := 0

	set := make(map[uint8]struct{})

	thisOneStart := 0
	for i := 0; i < len(s); i++ {
		if _, ok := set[s[i]]; !ok {
			set[s[i]] = struct{}{}
			continue
		}

		// 有重复

		// 是否更新最大值
		if (i - thisOneStart) > maxLen {
			maxLen = i - thisOneStart
		}

		// 开始下一轮
		for j := thisOneStart; ; j++ {
			if s[j] == s[i] {
				delete(set, s[j])
				thisOneStart = j + 1
				break
			}
			delete(set, s[j])
		}

		// warning: 开启下一轮之前加回去(出错点!!!!)
		set[s[i]] = struct{}{}
	}

	if len(s) - thisOneStart > maxLen {
		maxLen = len(s) - thisOneStart
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
