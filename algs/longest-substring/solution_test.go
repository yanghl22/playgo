package algs

import "testing"

type question1 struct {
	inp string
	otp int
}

func Test_LengthOfLongestSubstring(t *testing.T) {

	list := []question1{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}}

	for _, val := range list {
		result := lengthOfLongestSubstring(val.inp)
		if result == val.otp {
			t.Logf("%s success, expected %d, got %d ", val.inp, val.otp, result)
		} else {
			t.Errorf("%s failed, expected %d, got %d ", val.inp, val.otp, result)
		}
	}
}

func Test_LengthOfLongestSubstring2(t *testing.T) {

	list := []question1{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}}

	for _, val := range list {
		result := lengthOfLongestSubstring2(val.inp)
		if result == val.otp {
			t.Logf("%s success, expected %d, got %d ", val.inp, val.otp, result)
		} else {
			t.Errorf("%s failed, expected %d, got %d ", val.inp, val.otp, result)
		}
	}
}
func Test_LengthOfLongestSubstring1(t *testing.T) {

	list := []question1{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}}

	for _, val := range list {
		result := lengthOfLongestSubstring1(val.inp)
		if result == val.otp {
			t.Logf("%s success, expected %d, got %d ", val.inp, val.otp, result)
		} else {
			t.Errorf("%s failed, expected %d, got %d ", val.inp, val.otp, result)
		}
	}
}
