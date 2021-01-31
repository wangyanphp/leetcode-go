package main

//
func longestPalindrome(s string) string {

	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]

}

func expand(s string, i, j int) (int, int) {

	for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
		i--
		j++
	}

	return i + 1, j - 1
}

