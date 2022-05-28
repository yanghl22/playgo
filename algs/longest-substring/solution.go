package algs

func lengthOfLongestSubstring(s string) int {
	left, right, result := 0, 0, 0
	bitSet := make([]bool, 127)
	for left+result < len(s) && right < len(s) {
		if bitSet[s[right]] == true {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		result = max(result, right-left)
	}
	return result
}

func lengthOfLongestSubstring1(s string) int {
	left, right, result := 0, 0, 0
	indexes := make(map[byte]int, len(s))

	for right < len(s) && left+result < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			left = idx + 1
		} else {
			indexes[s[right]] = right
			right++
		}

		result = max(result, right-left)
	}
	return result

}

func lengthOfLongestSubstring2(s string) int {
	freqs := [127]int{}
	left, right, result := 0, 0, 0
	for right < len(s) && left+result < len(s) {
		if freqs[s[right]] != 0 {
			freqs[s[left]]--
			left++
		} else {
			freqs[s[right]]++
			right++
		}

		result = max(result, right-left)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
